package azure

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-azure/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeDiskMetricReadOpsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_metric_read_ops_hourly",
		Description: "Azure Compute Disk Metrics - Read Ops (Hourly)",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeDiskReadOpsHourly,
		},
		Columns: opengovernanceMonitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the disk.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.MonitoringMetric.DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}
