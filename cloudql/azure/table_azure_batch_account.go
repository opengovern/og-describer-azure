package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAzureBatchAccount(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_batch_account",
		Description: "Azure Batch Account",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetBatchAccount,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceNotFound", "ResourceGroupNotFound", "Invalid input"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListBatchAccount,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.Subscription"),
			},
			{
				Name:        "name",
				Description: "The resource name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "id",
				Description: "The resource identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.ID")},
			{
				Name:        "type",
				Description: "The resource type.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioned state of the batch account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.ProvisioningState")},
			{
				Name:        "account_endpoint",
				Description: "The account endpoint used to interact with the batch service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.AccountEndpoint")},
			{
				Name:        "active_job_and_job_schedule_quota",
				Description: "Active job and job schedule quota of the batch account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.ActiveJobAndJobScheduleQuota")},
			{
				Name:        "dedicated_core_quota",
				Description: "The dedicated core quota of the batch account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Account.Properties.DedicatedCoreQuota")},
			{
				Name:        "dedicated_core_quota_per_vm_family_enforced",
				Description: "Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not yet be enforced.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.Account.Properties.DedicatedCoreQuotaPerVMFamilyEnforced")},
			{
				Name:        "low_priority_core_quota",
				Description: "The low priority core quota of the batch account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Account.Properties.LowPriorityCoreQuota")},
			{
				Name:        "pool_allocation_mode",
				Description: "The pool allocation mode of the batch account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Properties.PoolAllocationMode")},
			{
				Name:        "pool_quota",
				Description: "The pool quota of the batch account.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.Account.Properties.PoolQuota")},
			{
				Name:        "public_network_access",
				Description: "Indicates whether or not public network access is allowed for the batch account.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "auto_storage",
				Description: "The auto storage properties of the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.AutoStorage")},
			{
				Name:        "dedicated_core_quota_per_vm_family",
				Description: "A list of the dedicated core quota per virtual machine family for the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.DedicatedCoreQuotaPerVMFamily")},
			{
				Name:        "diagnostic_settings",
				Description: "A list of active diagnostic settings for the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.DiagnosticSettingsResources")},
			{
				Name:        "encryption",
				Description: "Properties to enable customer managed key for the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.Encryption")},
			{
				Name:        "identity",
				Description: "The identity of the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Identity")},
			{
				Name:        "key_vault_reference",
				Description: "Key vault reference of the batch account.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Properties.KeyVaultReference")},
			{
				Name:        "private_endpoint_connections",
				Description: "The properties associated with the private endpoint connection.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.Account.Properties.PrivateEndpointConnections")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Account.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Account.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.Account.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.Account.Location").Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				// Check if context has been cancelled or if the limit has been hit (if specified)
				// if there is a limit, it will return the number of rows required to reach this limit
				Transform: transform.

					// Check if context has been cancelled or if the limit has been hit (if specified)
					// if there is a limit, it will return the number of rows required to reach this limit
					FromField("Description.ResourceGroup")},
		}),
	}
}

//// HYDRATE FUNCTIONS

// Create session

// Return nil, if no input provide

// Create session

// If we return the API response directly, the output only gives
// the contents of DiagnosticSettings
