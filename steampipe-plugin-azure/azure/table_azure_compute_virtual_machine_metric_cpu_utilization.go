package azure

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-azure/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeVirtualMachineMetricCpuUtilization(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_virtual_machine_metric_cpu_utilization",
		Description: "Azure Compute Virtual Machine Metrics - CPU Utilization",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListComputeVirtualMachineCpuUtilization,
		},
		Columns: opengovernanceMonitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the virtual machine.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.MonitoringMetric.DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}
