package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureVirtualMachineImagesImageTemplates(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_virtualmachineimages_imagetemplates",
		Description: "Azure VirtualMachineImages ImageTemplates",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetVirtualMachineImagesImageTemplates,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListVirtualMachineImagesImageTemplates,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the imagetemplates.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageTemplates.ID")},
			{
				Name:        "name",
				Description: "The name of the imagetemplates.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageTemplate.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ImageTemplate.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.ImageTemplate.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.ImageTemplate.ID").Transform(idToAkas),
			},
		}),
	}
}
