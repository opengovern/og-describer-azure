//go:generate go run ../SDK/runable/resourceType/resource_types_generator.go  --output resource_types.go --index-map ../steampipe/table_index_map.go && gofmt -w -s resource_types.go  && goimports -w resource_types.go

package describer

import (
	"context"
	"fmt"
	model "github.com/opengovern/og-describer-azure/pkg/SDK/models"
	"github.com/opengovern/og-describer-azure/provider"
	"github.com/opengovern/og-describer-azure/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"go.uber.org/zap"
	"sort"
	"strings"
)

func ListResourceTypes() []string {
	var list []string
	for k := range provider.ResourceTypes {
		list = append(list, k)
	}

	sort.Strings(list)
	return list
}

func ListSummarizeResourceTypes() []string {
	var list []string
	for k, v := range provider.ResourceTypes {
		if v.Summarize {
			list = append(list, k)
		}
	}

	sort.Strings(list)
	return list
}

func GetResourceType(resourceType string) (*model.ResourceType, error) {
	if r, ok := provider.ResourceTypes[resourceType]; ok {
		return &r, nil
	}
	resourceType = strings.ToLower(resourceType)
	for k, v := range provider.ResourceTypes {
		k := strings.ToLower(k)
		v := v
		if k == resourceType {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("resource type %s not found", resourceType)
}

func GetResourceTypesMap() map[string]model.ResourceType {
	return provider.ResourceTypes
}

func GetResourceTypeByTerraform(terraformType string) string {
	for t, v := range provider.ResourceTypes {
		for _, name := range v.TerraformName {
			if name == terraformType {
				return t
			}
		}
	}
	return ""
}

func GetResources(
	ctx context.Context,
	logger *zap.Logger,
	resourceType string,
	triggerType enums.DescribeTriggerType,
	cfg provider.AccountConfig,
	additionalData map[string]string,
	stream *model.StreamSender,
) error {
	_, err := describe(ctx, logger, cfg, resourceType, triggerType, additionalData, stream)
	if err != nil {
		return err
	}
	return nil
}

func describe(ctx context.Context, logger *zap.Logger, accountCfg provider.AccountConfig, resourceType string, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
	resourceTypeObject, ok := provider.ResourceTypes[resourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported resource type: %s", resourceType)
	}
	ctx = describer.WithLogger(ctx, logger)

	return resourceTypeObject.ListDescriber(ctx, accountCfg, triggerType, additionalData, stream)
}