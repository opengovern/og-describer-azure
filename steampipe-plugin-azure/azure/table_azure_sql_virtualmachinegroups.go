package azure

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-azure/pkg/sdk/generated"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureSqlVirtualMachineGroups(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_sql_virtualmachinegroups",
		Description: "Azure Sql VirtualMachineGroups",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetSqlServerVirtualMachineGroup,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSqlServerVirtualMachineGroup,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the virtualmachinegroups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.ID")},
			{
				Name:        "name",
				Description: "The name of the virtualmachinegroups.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Group.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.Group.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.Group.ID").Transform(idToAkas),
			},
		}),
	}
}
