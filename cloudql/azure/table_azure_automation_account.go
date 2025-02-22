package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/automation/mgmt/2019-06-01/automation"
	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION ////

func tableAzureApAutomationAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_automation_account",
		Description: "Azure Automation Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetAutomationAccounts,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAutomationAccounts,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.Subscription"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the resource.",
				Transform:   transform.FromField("Description.Automation.Name")},
			{
				Name:        "id",
				Description: "Fully qualified resource ID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.ID")},
			{
				Name:        "description",
				Description: "The description for the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.Description")},
			{
				Name:        "etag",
				Description: "Gets the etag of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Etag")},
			{
				Name:        "creation_time",
				Description: "The creation time of the account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Automation.Properties.CreationTime")},
			{
				Name:        "last_modified_time",
				Description: "The last modified time of the account.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.Automation.Properties.LastModifiedTime")},
			{
				Name:        "last_modified_by",
				Description: "The account last modified by.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.LastModifiedBy")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Type")},
			{
				Name:        "state",
				Description: "The status of account. Possible values include: 'AccountStateOk', 'AccountStateUnavailable', 'AccountStateSuspended'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.State")},
			{
				Name:        "sku_name",
				Description: "The SKU name of the account. Possible values include: 'SkuNameEnumFree', 'SkuNameEnumBasic'.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.SKU.Name")},
			{
				Name:        "sku_family",
				Description: "The SKU family of the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.SKU.Family")},
			{
				Name:        "sku_capacity",
				Description: "The SKU capacity of the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Properties.SKU.Capacity")},

			// Steampipe standard columns
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Automation.ID").Transform(idToAkas),
			},

			// Azure standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Automation.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceGroup").Transform(toLower),
			},
		}),
	}
}

//// LIST FUNCTION ////

func listAutomationAccounts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	session, err := GetNewSession(ctx, d, "MANAGEMENT")
	if err != nil {
		plugin.Logger(ctx).Error("azure_automation_variable.listAutomationAccounts", "session_error", err)
		return nil, err
	}
	subscriptionID := session.SubscriptionID

	accountClient := automation.NewAccountClientWithBaseURI(session.ResourceManagerEndpoint, subscriptionID)
	accountClient.Authorizer = session.Authorizer

	result, err := accountClient.List(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("azure_automation_variable.listAutomationAccounts", "api_error", err)
		return nil, err
	}

	for _, account := range result.Values() {
		d.StreamListItem(ctx, account)
		// Check if context has been cancelled or if the limit has been hit (if specified)
		// if there is a limit, it will return the number of rows required to reach this limit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	for result.NotDone() {
		err = result.NextWithContext(ctx)
		if err != nil {
			plugin.Logger(ctx).Error("azure_automation_variable.listAutomationAccounts", "paginator_error", err)
			return nil, err
		}

		for _, account := range result.Values() {
			d.StreamListItem(ctx, account)
			// Check if context has been cancelled or if the limit has been hit (if specified)
			// if there is a limit, it will return the number of rows required to reach this limit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}
	return nil, err
}

//// HYDRATE FUNCTIONS ////

func getAutomationAccount(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	name := d.EqualsQuals["name"].GetStringValue()
	resourceGroup := d.EqualsQuals["resource_group"].GetStringValue()

	session, err := GetNewSession(ctx, d, "MANAGEMENT")
	if err != nil {
		plugin.Logger(ctx).Error("azure_automation_variable.getAutomationAccount", "session_error", err)
		return nil, err
	}
	subscriptionID := session.SubscriptionID

	accountClient := automation.NewAccountClientWithBaseURI(session.ResourceManagerEndpoint, subscriptionID)
	accountClient.Authorizer = session.Authorizer

	op, err := accountClient.Get(ctx, resourceGroup, name)
	if err != nil {
		plugin.Logger(ctx).Error("azure_automation_variable.getAutomationAccount", "api_error", err)
		return nil, err
	}

	// In some cases resource does not give any notFound error
	// Instead it returns empty data
	if op.ID != nil {
		return op, nil
	}

	return nil, nil
}
