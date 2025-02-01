package describers

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/opengovern/og-describer-azure/discovery/pkg/models"
	model "github.com/opengovern/og-describer-azure/discovery/provider"
)

func CdnProfiles(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcdn.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}
	client := clientFactory.NewProfilesClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getCdnProfiles(ctx, v, subscription)
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

func getCdnProfiles(ctx context.Context, v *armcdn.Profile, subscription string) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.CDNProfileDescription{
			Profile:       *v,
			ResourceGroup: resourceGroup,
			Subscription:  subscription,
		},
	}
	return &resource
}

func CdnEndpoint(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	clientFactory, err := armcdn.NewClientFactory(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	client := clientFactory.NewProfilesClient()
	endpointsClient := clientFactory.NewEndpointsClient()

	var values []models.Resource
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resources, err := getCdnProfilesEndpoints(ctx, endpointsClient, v, subscription)
			if err != nil {
				return nil, err
			}
			for _, resource := range resources {
				if stream != nil {
					if err := (*stream)(resource); err != nil {
						return nil, err
					}
				} else {
					values = append(values, resource)
				}
			}
		}
	}
	return values, nil
}

func getCdnProfilesEndpoints(ctx context.Context, endpointsClient *armcdn.EndpointsClient, v *armcdn.Profile, subscription string) ([]models.Resource, error) {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	pager := endpointsClient.NewListByProfilePager(resourceGroup, *v.Name, nil)

	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, endpoint := range page.Value {
			resource := getCdnEndpoint(ctx, v, subscription, endpoint, resourceGroup)
			values = append(values, *resource)
		}
	}
	return values, nil
}

func getCdnEndpoint(ctx context.Context, v *armcdn.Profile, subscription string, endpoint *armcdn.Endpoint, resourceGroup string) *models.Resource {
	return &models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.CDNEndpointDescription{
			Endpoint:      *endpoint,
			ResourceGroup: resourceGroup,
			Subscription:  subscription,
		},
	}
}
