package provider

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	model "github.com/opengovern/og-describer-azure/pkg/SDK/models"
	"github.com/opengovern/og-describer-azure/provider/configs"

	"github.com/opengovern/og-describer-azure/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
)

func DescribeBySubscription(describe func(context.Context, *azidentity.ClientSecretCredential, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.AccountCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientSecret, nil)
		if err != nil {
			return nil, err
		}
		var values []model.Resource
		result, err := describe(ctx, cred, additionalData["subscriptionId"], stream)
		if err != nil {
			return nil, err
		}
		accountInfo := map[string]string{
			"SubscriptionID": additionalData["subscriptionId"],
			"TenantID":       cfg.TenantID,
		}
		for _, resource := range result {
			resource.AccountInfo = accountInfo
		}
		values = append(values, result...)

		return values, nil
	}
}

func DescribeADByTenantID(describe func(context.Context, *azidentity.ClientSecretCredential, string, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.AccountCredentials, triggerType enums.DescribeTriggerType, additionalData map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)
		cred, err := azidentity.NewClientSecretCredential(cfg.TenantID, cfg.ClientID, cfg.ClientSecret, nil)
		if err != nil {
			return nil, err
		}
		var values []model.Resource
		result, err := describe(ctx, cred, cfg.TenantID, stream)
		if err != nil {
			return nil, err
		}

		values = append(values, result...)

		return values, nil
	}
}
