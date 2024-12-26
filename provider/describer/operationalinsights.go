package describer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"github.com/opengovern/og-describer-azure/pkg/sdk/models"
	"strings"

	"github.com/opengovern/og-describer-azure/provider/model"
)

func OperationalInsightsWorkspaces(ctx context.Context, cred *azidentity.ClientSecretCredential, subscription string, stream *models.StreamSender) ([]models.Resource, error) {
	client, err := armoperationalinsights.NewWorkspacesClient(subscription, cred, nil)
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(nil)
	var values []models.Resource
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range page.Value {
			resource := getOperationalInsightsWorkspace(ctx, v)
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

func getOperationalInsightsWorkspace(ctx context.Context, v *armoperationalinsights.Workspace) *models.Resource {
	resourceGroup := strings.Split(*v.ID, "/")[4]

	resource := models.Resource{
		ID:       *v.ID,
		Name:     *v.Name,
		Location: *v.Location,
		Description: model.OperationalInsightsWorkspacesDescription{
			Workspace:     *v,
			ResourceGroup: resourceGroup,
		},
	}
	return &resource
}
