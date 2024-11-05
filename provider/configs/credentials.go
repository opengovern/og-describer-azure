package configs

import (
	"encoding/json"
	"fmt"
	model "github.com/opengovern/og-describer-azure/pkg/SDK/models"
	azuremodel "github.com/opengovern/og-describer-azure/provider/model"
	"github.com/opengovern/og-util/pkg/describe"
	"strings"
)

type AccountCredentials struct {
	TenantID        string `json:"azure_tenant_id"`
	ClientID        string `json:"azure_client_id"`
	ClientSecret    string `json:"azure_client_password"`
	CertificatePath string `json:"certificatePath"`
	CertificatePass string `json:"certificatePass"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

func AccountConfigFromMap(m map[string]any) (AccountCredentials, error) {
	mj, err := json.Marshal(m)
	if err != nil {
		return AccountCredentials{}, err
	}

	var c AccountCredentials
	err = json.Unmarshal(mj, &c)
	if err != nil {
		return AccountCredentials{}, err
	}

	return c, nil
}

func GetResourceMetadata(job describe.DescribeJob, resource model.Resource) (map[string]string, error) {
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
		return nil, fmt.Errorf("marshal metadata: %v", err.Error())
	}

	metadata := make(map[string]string)
	err = json.Unmarshal(azureMetadataBytes, &metadata)
	if err != nil {
		return nil, fmt.Errorf("unmarshal metadata: %v", err.Error())
	}
	return metadata, nil
}

func AdjustResource(job describe.DescribeJob, resource *model.Resource) error {
	resource.Location = fixAzureLocation(resource.Location)
	resource.Type = strings.ToLower(job.ResourceType)
	return nil
}

func fixAzureLocation(l string) string {
	return strings.ToLower(strings.ReplaceAll(l, " ", ""))
}

func GetAdditionalParameters(job describe.DescribeJob) (map[string]string, error) {
	additionalParameters := make(map[string]string)
	additionalParameters["subscriptionId"] = job.AccountID

	return additionalParameters, nil
}
