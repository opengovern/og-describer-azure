package azure

import (
	"context"

	opengovernance "github.com/opengovern/og-describer-azure/discovery/pkg/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureRedisCache(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_redis_cache",
		Description: "Azure Redis Cache",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"name", "resource_group"}),
			Hydrate:    opengovernance.GetRedisCache,
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound", "ResourceNotFound", "400", "400"}),
			},
		},
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListRedisCache,
		},
		Columns: azureOGColumns([]*plugin.Column{
			{
				Name:        "subscription",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Subscription"),
			},
			{
				Name:        "name",
				Description: "The name of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Name")},
			{
				Name:        "id",
				Description: "The unique id identifying the resource in subscription.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.ID")},
			{
				Name:        "type",
				Description: "The type of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Type")},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the redis instance at the time the operation was called. Valid values are: 'Creating', 'Deleting', 'Disabled', 'Failed', 'Linking', 'Provisioning', 'RecoveringScaleFailure', 'Scaling', 'Succeeded', 'Unlinking', 'Unprovisioning', and 'Updating'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Properties.ProvisioningState"),
			},
			{
				Name:        "redis_version",
				Description: "Specifies the version.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.RedisVersion")},
			{
				Name:        "enable_non_ssl_port",
				Description: "Specifies whether the non-ssl Redis server port (6379) is enabled.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.EnableNonSSLPort")},
			{
				Name:        "host_name",
				Description: "Specifies the name of the redis host.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.HostName")},
			{
				Name:        "minimum_tls_version",
				Description: "Specifies the TLS version requires to connect.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Properties.MinimumTLSVersion").Transform(transform.NullIfZeroValue),
			},
			{
				Name:        "port",
				Description: "Specifies the redis non-SSL port.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Port")},
			{
				Name:        "public_network_access",
				Description: "Indicates whether or not public endpoint access is allowed for this cache. Valid values are: 'Enabled', 'Disabled'.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Properties.PublicNetworkAccess"),
			},
			{
				Name:        "sku_capacity",
				Description: "The size of the Redis cache to deploy.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.SKU.Capacity")},
			{
				Name:        "sku_family",
				Description: "The SKU family to use.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Properties.SKU.Family"),
			},
			{
				Name:        "sku_name",
				Description: "The type of Redis cache to deploy.",
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Properties.SKU.Name"),
			},
			{
				Name:        "ssl_port",
				Description: "Specifies the redis SSL port.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.SSLPort")},
			{
				Name:        "subnet_id",
				Description: "The full resource ID of a subnet in a virtual network to deploy the Redis cache in.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.SubnetID")},
			{
				Name:        "static_ip",
				Description: "Specifies the statis IP address. Required when deploying a Redis cache inside an existing Azure Virtual Network.",
				Type:        proto.ColumnType_IPADDR,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.StaticIP")},
			{
				Name:        "replicas_per_master",
				Description: "The number of replicas to be created per master.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.ReplicasPerMaster")},
			{
				Name:        "shard_count",
				Description: "The number of shards to be created on a premium cluster cache.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.ShardCount")},
			{
				Name:        "access_keys",
				Description: "The keys of the Redis cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.AccessKeys")},
			{
				Name:        "linked_servers",
				Description: "A list of the linked servers associated with the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.LinkedServers")},
			{
				Name:        "instances",
				Description: "A list of the Redis instances associated with the cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.Instances")},
			{
				Name:        "private_endpoint_connections",
				Description: "A list of private endpoint connection associated with the specified redis cache.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.PrivateEndpointConnections")},
			{
				Name:        "redis_configuration",
				Description: "Describes the redis cache configuration.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.RedisConfiguration")},
			{
				Name:        "tenant_settings",
				Description: "A dictionary of tenant settings.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Properties.TenantSettings")},
			{
				Name:        "zones",
				Description: "A list of availability zones denoting where the resource needs to come from.",
				Type:        proto.ColumnType_JSON,

				// Steampipe standard columns
				Transform: transform.FromField("Description.ResourceInfo.Zones")},

			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ResourceInfo.Name")},
			{
				Name:        "tags",
				Description: ColumnDescriptionTags,
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.ResourceInfo.Tags")},
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,

				// Azure standard columns

				Transform: transform.FromField("Description.ResourceInfo.ID").Transform(idToAkas),
			},

			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,

				Transform: transform.FromField("Description.ResourceInfo.Location").Transform(formatRegion).Transform(toLower),
			},
			{
				Name:        "resource_group",
				Description: ColumnDescriptionResourceGroup,
				Type:        proto.ColumnType_STRING,

				//// LIST FUNCTION

				Transform: transform.

					// Create session
					FromField("Description.ResourceGroup")},
		}),
	}
}

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

// Check if context has been cancelled or if the limit has been hit (if specified)
// if there is a limit, it will return the number of rows required to reach this limit

//// HYDRATE FUNCTIONS

// Return nil, if no input provided

// Create session
