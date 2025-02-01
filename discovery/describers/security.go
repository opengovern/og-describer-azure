package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func SecurityCenterAutoProvisioning(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAutoProvisioningSettingsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterAutoProvisioning(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterAutoProvisioning(ctx context.Context, v *armsecurity.AutoProvisioningSetting, subscription string) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.SecurityCenterAutoProvisioningDescription{
			AutoProvisioningSetting: *v,
			Subscription:            subscription,
		},
	}

	return &resource
}

func SecurityCenterContact(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewContactsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.ContactList.Value {
			resource := GetSecurityCenterContact(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterContact(ctx context.Context, v *armsecurity.Contact, subscription string) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.SecurityCenterContactDescription{
			Contact:      *v,
			Subscription: subscription,
		},
	}
	return &resource
}

func SecurityCenterJitNetworkAccessPolicy(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewJitNetworkAccessPoliciesClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterJitNetworkAccessPolicy(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterJitNetworkAccessPolicy(ctx context.Context, v *armsecurity.JitNetworkAccessPolicy, subscription string) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SecurityCenterJitNetworkAccessPolicyDescription{
			JitNetworkAccessPolicy: *v,
			Subscription:           subscription,
		},
	}
	return &resource
}

func SecurityCenterSetting(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSettingsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterSetting(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterSetting(ctx context.Context, v armsecurity.SettingClassification, subscription string) *models.Resource {
	var settingStatus bool
	if *v.GetSetting().Kind == armsecurity.SettingKindDataExportSettings {
		settingStatus = true
	} else {
		settingStatus = false
	}
	resource := models.Resource{
		ID:       *v.GetSetting().ID,
		Name:     *v.GetSetting().Name,
		Location: "global",
		Description: model.SecurityCenterSettingDescription{
			Setting:             *v.GetSetting(),
			ExportSettingStatus: settingStatus,
			Subscription:        subscription,
		},
	}
	return &resource
}

func SecurityCenterSubscriptionPricing(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewPricingsClient()

	var values []models.Resource
	list, err := client.List(ctx, "", nil)
	if err != nil {
		return nil, err
	}
	for _, v := range list.Value {
		resource := GetSecurityCenterSubscriptionPricing(ctx, v, subscription)
		if stream != nil {
			if err := (*stream)(*resource); err != nil {
				return nil, err
			}
		} else {
			values = append(values, *resource)
		}
	}
	return values, nil
}

func GetSecurityCenterSubscriptionPricing(ctx context.Context, v *armsecurity.Pricing, subscription string) *models.Resource {
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: "global",
		Description: model.SecurityCenterSubscriptionPricingDescription{
			Pricing:      *v,
			Subscription: subscription,
		},
	}
	return &resource
}

func SecurityCenterAutomation(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewAutomationsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterAutomation(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterAutomation(ctx context.Context, v *armsecurity.Automation, subscription string) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]
	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.SecurityCenterAutomationDescription{
			Automation:    *v,
			ResourceGroup: resourceGroup,
			Subscription:  subscription,
		},
	}
	return &resource
}

func SecurityCenterSubAssessment(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armsecurity.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewSubAssessmentsClient()

	var values []models.Resource
	pager := client.NewListAllPager("subscriptions/"+subscription, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := GetSecurityCenterSubAssessment(ctx, v, subscription)
			if stream != nil {
				if err := (*stream)(*resource); err != nil {
					return nil, err
				}
			} else {
				values = append(values, *resource)
			}
		}
	}
	return values, nil
}

func GetSecurityCenterSubAssessment(ctx context.Context, v *armsecurity.SubAssessment, subscription string) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Location: "global",
		Name:     *v.Name,
		Description: model.SecurityCenterSubAssessmentDescription{
			SubAssessment: *v,
			ResourceGroup: resourceGroup,
			Subscription:  subscription,
		},
	}
	return &resource
}
