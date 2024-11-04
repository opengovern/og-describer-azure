package azure

import (
	"context"
	"github.com/opengovern/og-azure-describer-new/SDK/generated"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureStreamAnalyticsCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_streamanalytics_cluster",
		Description: "Azure StreamAnalytics Cluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetStreamAnalyticsCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListStreamAnalyticsCluster,
		},
		Columns: azureKaytuColumns([]*plugin.Column{
			{
				Name:        "id",
				Description: "The id of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Cluster.ID")},
			{
				Name:        "name",
				Description: "The name of the cluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.SKU.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.StreamingJob.SKU.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.StreamingJob.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.StreamingJob.Properties.ClusterID").Transform(idToAkas),
			},
		}),
	}
}
