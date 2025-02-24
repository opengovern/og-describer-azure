package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureIamRoleDefinition(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_role_definition",
		Description: "Azure Role Definition",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    opengovernance.GetRoleDefinition,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoleDefinition,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "short_id",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly ID/name that identifies the role definition.",
				Transform:   transform.FromField("Description.RoleDefinition.Name")},
			{
				Name:        "id",
				Description: "Contains GUID to identify a role definition uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.ID")},
			{
				Name:        "type",
				Description: "Contains the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.Type")},
			{
				Name:        "name",
				Description: "Name of the Role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.RoleName")},
			{
				Name:        "type",
				Description: "type of the role definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.RoleType")},
			{
				Name:        "description",
				Description: "Description of the role definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.Description")},
			{
				Name:        "assignable_scopes",
				Description: "A list of assignable scopes for which the role definition can be assigned.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.AssignableScopes")},
			{
				Name:        "permissions",
				Description: "A list of actions, which can be accessed.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.Permissions")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleDefinition.Properties.RoleName").Transform(toLower)},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				//// LIST FUNCTION

				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.RoleDefinition.ID").Transform(idToAkas),
			},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS
