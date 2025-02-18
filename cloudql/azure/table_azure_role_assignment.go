package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureIamRoleAssignment(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_role_assignment",
		Description: "Azure Role Assignment",
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetRoleAssignment,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRoleAssignment,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:      "subscription",
				Type:      proto.ColumnType_STRING,
				Transform: transform.FromField("Description.Subscription"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The friendly name that identifies the role assignment.",
				Transform:   transform.FromField("Description.RoleAssignment.name")},
			{
				Name:        "id",
				Description: "Contains ID to identify a role assignment uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.id")},
			{
				Name:        "scope",
				Description: "Current state of the role assignment.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.properties.scope")},
			{
				Name:        "type",
				Description: "Contains the resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.type")},
			{
				Name:        "principal_id",
				Description: "Contains the principal id.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.properties.principalId")},
			{
				Name:        "principal_type",
				Description: "Principal type of the assigned principal ID.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.RoleAssignment.properties.principalType"),
			},
			{
				Name:        "created_on",
				Description: "Time it was created.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RoleAssignment.properties.createdOn"),
			},
			{
				Name:        "updated_on",
				Description: "Time it was updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.RoleAssignment.properties.updatedOn"),
			},
			{
				Name:        "role_definition_id",
				Description: "Name of the assigned role definition.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.properties.roleDefinitionID")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.RoleAssignment.name")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.RoleAssignment.id").Transform(idToAkas),
			},
		}),
	}
}
