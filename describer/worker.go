package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-azure-describer-new/provider"
	azuremodel "github.com/opengovern/og-azure-describer-new/provider/model"
	"github.com/opengovern/og-util/pkg/es"
	"github.com/opengovern/og-util/pkg/integration"
	"strings"

	"github.com/go-errors/errors"
	"github.com/opengovern/og-azure-describer-new/provider/describer"
	"github.com/opengovern/og-azure-describer-new/steampipe"
	"github.com/opengovern/og-util/pkg/describe"
	"github.com/opengovern/og-util/pkg/source"
	"github.com/opengovern/og-util/pkg/vault"
	"go.uber.org/zap"
)

type Error struct {
	ErrCode string

	error
}

func fixAzureLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", ""))
}

func trimEmptyMaps(input map[string]any) {
	for key, value := range input {
		switch value.(type) {
		case map[string]any:
			if len(value.(map[string]any)) != 0 {
				trimEmptyMaps(value.(map[string]any))
			}
			if len(value.(map[string]any)) == 0 {
				delete(input, key)
			}
		}
	}
}

func trimJsonFromEmptyObjects(input []byte) ([]byte, error) {
	unknownData := map[string]any{}
	err := json.Unmarshal(input, &unknownData)
	if err != nil {
		return nil, err
	}
	trimEmptyMaps(unknownData)
	return json.Marshal(unknownData)
}

func Do(ctx context.Context,
	vlt vault.VaultSourceConfig,
	logger *zap.Logger,
	job describe.DescribeJob,
	grpcEndpoint string,
	describeDeliverToken string,
	ingestionPipelineEndpoint string,
	useOpenSearch bool) (resourceIDs []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("paniced with error: %v", r)
			logger.Error("paniced with error", zap.Error(err), zap.String("stackTrace", errors.Wrap(r, 2).ErrorStack()))
		}
	}()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	config, err := vlt.Decrypt(ctx, job.CipherText)
	if err != nil {
		return nil, fmt.Errorf("decrypt error: %w", err)
	}
	logger.Info("decrypted config", zap.Any("config", config))

	return doDescribe(ctx, logger, job, config, grpcEndpoint, ingestionPipelineEndpoint, describeDeliverToken, useOpenSearch)
}

func doDescribe(
	ctx context.Context,
	logger *zap.Logger,
	job describe.DescribeJob,
	config map[string]any,
	grpcEndpoint, ingestionPipelineEndpoint string,
	describeToken string,
	useOpenSearch bool) ([]string, error) {
	rs, err := NewResourceSender(grpcEndpoint, ingestionPipelineEndpoint, describeToken, job.JobID, useOpenSearch, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to resource sender: %w", err)
	}

	plg := steampipe.Plugin()
	plgAD := steampipe.ADPlugin()
	creds, err := provider.AccountConfigFromMap(config)
	if err != nil {
		return nil, fmt.Errorf("azure subscription credentials: %w", err)
	}
	subscriptionId := job.AccountID

	f := func(resource describer.Resource) error {
		if resource.Description == nil {
			return nil
		}
		descriptionJSON, err := json.Marshal(resource.Description)
		if err != nil {
			return fmt.Errorf("failed to marshal description: %w", err)
		}
		descriptionJSON, err = trimJsonFromEmptyObjects(descriptionJSON)
		if err != nil {
			return fmt.Errorf("failed to trim json: %w", err)
		}
		resource.Location = fixAzureLocation(resource.Location)
		resource.SubscriptionID = subscriptionId
		resource.Type = strings.ToLower(job.ResourceType)
		azureMetadata := azuremodel.Metadata{
			ID:               resource.ID,
			Name:             resource.Name,
			SubscriptionID:   job.AccountID,
			Location:         resource.Location,
			CloudEnvironment: "AzurePublicCloud",
			ResourceType:     strings.ToLower(job.ResourceType),
			SourceID:         job.SourceID,
		}
		azureMetadataBytes, err := json.Marshal(azureMetadata)
		if err != nil {
			return fmt.Errorf("marshal metadata: %v", err.Error())
		}

		metadata := make(map[string]string)
		err = json.Unmarshal(azureMetadataBytes, &metadata)
		if err != nil {
			return fmt.Errorf("unmarshal metadata: %v", err.Error())
		}

		desc := resource.Description
		err = json.Unmarshal(descriptionJSON, &desc)
		if err != nil {
			return fmt.Errorf("unmarshal description: %v", err.Error())
		}

		kafkaResource := Resource{
			ID:              resource.UniqueID(),
			Name:            resource.Name,
			ResourceGroup:   resource.ResourceGroup,
			Location:        resource.Location,
			IntegrationType: source.CloudAzure,
			ResourceType:    strings.ToLower(job.ResourceType),
			ResourceJobID:   job.JobID,
			SourceID:        job.SourceID,
			CreatedAt:       job.DescribedAt,
			Description:     desc,
			Metadata:        metadata,
		}

		tags, name, err := steampipe.ExtractTagsAndNames(logger, plg, plgAD, job.ResourceType, kafkaResource)
		if err != nil {
			logger.Error("failed to build tags for service", zap.Error(err), zap.String("resourceType", job.ResourceType), zap.Any("resource", kafkaResource))
			return fmt.Errorf("failed to build tags for servicem resource type: %v, resource: %v, err: %v", job.ResourceType, kafkaResource, err)
		}
		if len(name) > 0 {
			kafkaResource.Metadata["name"] = name
		}

		var description any
		err = json.Unmarshal([]byte(descriptionJSON), &description)
		if err != nil {
			logger.Error("failed to parse resource description json", zap.Error(err))
			return fmt.Errorf("failed to parse resource description json")
		}

		newTags := make([]es.Tag, 0, len(tags))
		for k, v := range tags {
			newTags = append(newTags, es.Tag{
				// tags should be case-insensitive
				Key:   strings.ToLower(k),
				Value: strings.ToLower(v),
			})
		}

		rs.Send(&es.Resource{
			ID:              resource.UniqueID(),
			Description:     description,
			IntegrationType: integration.Type("AZURE_SUBSCRIPTION"),
			ResourceType:    strings.ToLower(job.ResourceType),
			ResourceJobID:   uint(uint32(job.JobID)),
			SourceID:        job.SourceID,
			Metadata:        metadata,
			Name:            resource.Name,
			ResourceGroup:   resource.ResourceGroup,
			Location:        resource.Location,
			CreatedAt:       job.DescribedAt,
			CanonicalTags:   newTags,
		})
		return nil
	}
	clientStream := (*describer.StreamSender)(&f)

	_, err = provider.GetResources(
		ctx,
		logger,
		job.ResourceType,
		job.TriggerType,
		[]string{subscriptionId},
		creds,
		string(provider.AuthEnv),
		"",
		clientStream,
	)
	if err != nil {
		return nil, err
	}

	rs.Finish()

	return rs.GetResourceIDs(), nil
}
