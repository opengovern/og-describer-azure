package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAzureCosmosdbCassandraCluster(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_cosmosdb_cassandra_cluster",
		Description: "Azure Cosmosdb CassandraCluster",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"), //TODO: change this to the primary key columns in model.go
			Hydrate:    opengovernance.GetCosmosdbCassandraCluster,
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListCosmosdbCassandraCluster,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CassandraCluster.Properties.Subscription"),
			},
			{
				Name:        "id",
				Description: "The id of the cassandracluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CassandraCluster.ID")},
			{
				Name:        "name",
				Description: "The name of the cassandracluster.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CassandraCluster.Name")},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.CassandraCluster.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				// probably needs a transform function
				Transform: transform.FromField("Description.CassandraCluster.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				// or generate it below (keep the Transform(arnToTurbotAkas) or use Transform(transform.EnsureStringArray))
				Transform: transform.FromField("Description.CassandraCluster.ID").Transform(idToAkas),
			},
		}),
	}
}
