package provider
import (
	"github.com/opengovern/og-describer-azure/provider/describer"
	"github.com/opengovern/og-describer-azure/provider/configs"
	model "github.com/opengovern/og-describer-azure/pkg/sdk/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"Microsoft.App/containerApps": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.App/containerApps",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppContainerApps),
		GetDescriber:         nil,
	},

	"Microsoft.Blueprint/blueprints": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Blueprint/blueprints",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Blueprint.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.BlueprintBlueprint),
		GetDescriber:         nil,
	},

	"Microsoft.Cdn/profiles": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Cdn/profiles",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/CDN%20Profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CdnProfiles),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/cloudServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/cloudServices",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Cloud%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeCloudServices),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerInstance/containerGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ContainerInstance/containerGroups",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ContainerInstanceContainerGroups),
		GetDescriber:         nil,
	},

	"Microsoft.DataMigration/services": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataMigration/services",
		Tags:                 map[string][]string{
            "category": {"Migration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Database%20Migration%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataMigrationServices),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupVaults": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupVaults",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Backup%20vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataProtectionBackupVaults),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupJobs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupJobs",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataProtectionBackupJobs),
		GetDescriber:         nil,
	},

	"Microsoft.DataProtection/backupVaults/backupPolicies": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataProtection/backupVaults/backupPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Backup%20vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataProtectionBackupVaultsBackupPolicies),
		GetDescriber:         nil,
	},

	"Microsoft.Logic/integrationAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Logic/integrationAccounts",
		Tags:                 map[string][]string{
            "category": {"Integration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Integration%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LogicIntegrationAccounts),
		GetDescriber:         nil,
	},

	"Microsoft.Network/bastionHosts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/bastionHosts",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Bastion.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkBastionHosts),
		GetDescriber:         nil,
	},

	"Microsoft.Network/connections": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/connections",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Hybrid%20Connection.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkConnections),
		GetDescriber:         nil,
	},

	"Microsoft.Network/firewallPolicies": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/firewallPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Firewall%20Policy.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.FirewallPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Network/localNetworkGateways": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/localNetworkGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Local%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LocalNetworkGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateLinkServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/privateLinkServices",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Private%20link%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PrivateLinkService),
		GetDescriber:         nil,
	},

	"Microsoft.Network/publicIPPrefixes": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/publicIPPrefixes",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Public%20IP%20Prefix.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PublicIPPrefix),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualHubs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualHubs",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Virtual%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkVirtualHubs),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualWans": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualWans",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20WAN.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkVirtualWans),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnGateways": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.VpnGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnGateways/vpnConnections": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnGateways/vpnConnections",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkVpnGatewaysVpnConnections),
		GetDescriber:         nil,
	},

	"Microsoft.Network/vpnSites": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/vpnSites",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkVpnGatewaysVpnSites),
		GetDescriber:         nil,
	},

	"Microsoft.OperationalInsights/workspaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.OperationalInsights/workspaces",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.OperationalInsightsWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.StreamAnalytics/cluster": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.StreamAnalytics/cluster",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Stream%20Analytics%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StreamAnalyticsCluster),
		GetDescriber:         nil,
	},

	"Microsoft.TimeSeriesInsights/environments": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.TimeSeriesInsights/environments",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Time%20Series%20Insights%20Environment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.TimeSeriesInsightsEnvironments),
		GetDescriber:         nil,
	},

	"Microsoft.VirtualMachineImages/imageTemplates": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.VirtualMachineImages/imageTemplates",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Image%20Template.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.VirtualMachineImagesImageTemplates),
		GetDescriber:         nil,
	},

	"Microsoft.Web/serverFarms": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/serverFarms",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Server%20Farm.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.WebServerFarms),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets/virtualMachines": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets/virtualMachines",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Scale%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineScaleSetVm),
		GetDescriber:         nil,
	},

	"Microsoft.Automation/automationAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Automation/automationAccounts",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Automation%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AutomationAccounts),
		GetDescriber:         nil,
	},

	"Microsoft.Automation/automationAccounts/variables": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Automation/automationAccounts/variables",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Automation%20Variable.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AutomationVariables),
		GetDescriber:         nil,
	},

	"Microsoft.Network/dnsZones": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/dnsZones",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Zone%20(Public).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DNSZones),
		GetDescriber:         nil,
	},

	"Microsoft.Databricks/workspaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Databricks/workspaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Databricks.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DatabricksWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateDnsZones": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/privateDnsZones",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Zone%20(Private).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PrivateDnsZones),
		GetDescriber:         nil,
	},

	"Microsoft.Network/privateEndpoints": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/privateEndpoints",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Private%20Endpoint.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PrivateEndpoints),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkWatchers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/networkWatchers",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Watcher.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkWatcher),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/resourceGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/resourceGroups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ResourceGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Web/staticSites": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/staticSites",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Static%20Web%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppServiceWebApp),
		GetDescriber:         nil,
	},

	"Microsoft.Web/sites/slots": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/sites/slots",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Static%20Web%20App.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppServiceWebAppSlot),
		GetDescriber:         nil,
	},

	"Microsoft.CognitiveServices/accounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.CognitiveServices/accounts",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Cognitive%20Services.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CognitiveAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/managedInstances": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/managedInstances",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Managed%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MssqlManagedInstance),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/virtualclusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/virtualclusters",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Database.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlVirtualClusters),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/managedInstances/databases": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/managedInstances/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Managed%20Instance.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MssqlManagedInstanceDatabases),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers/databases": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Database.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/largeFileSharesState": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/largeFileSharesState",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/File%20Share.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageFileShare),
		GetDescriber:         nil,
	},

	"Microsoft.DBforPostgreSQL/servers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforPostgreSQL/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20PostgreSQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PostgresqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforPostgreSQL/flexibleservers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforPostgreSQL/flexibleservers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20PostgreSQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PostgresqlFlexibleservers),
		GetDescriber:         nil,
	},

	"Microsoft.AnalysisServices/servers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.AnalysisServices/servers",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Analysis%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AnalysisService),
		GetDescriber:         nil,
	},

	"Microsoft.Security/pricings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/pricings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterSubscriptionPricing),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/guestDiagnosticSettings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Insights/guestDiagnosticSettings",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Diagnostics%20Setting.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DiagnosticSetting),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/autoscaleSettings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Insights/autoscaleSettings",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AutoscaleSetting),
		GetDescriber:         nil,
	},

	"Microsoft.Web/hostingEnvironments": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/hostingEnvironments",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service%20Environment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppServiceEnvironment),
		GetDescriber:         nil,
	},

	"Microsoft.Cache/redis": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Cache/redis",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cache%20for%20Redis.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RedisCache),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerRegistry/registries": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ContainerRegistry/registries",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Container%20Registry.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ContainerRegistry),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories/pipelines": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories/pipelines",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataFactoryPipeline),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/resourceSku": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/resourceSku",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeResourceSKU),
		GetDescriber:         nil,
	},

	"Microsoft.Network/expressRouteCircuits": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/expressRouteCircuits",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/ExpressRoute%20Circuit.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ExpressRouteCircuit),
		GetDescriber:         nil,
	},

	"Microsoft.Management/managementgroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Management/managementgroups",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Management%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ManagementGroup),
		GetDescriber:         nil,
	},

	"microsoft.SqlVirtualMachine/SqlVirtualMachines": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "microsoft.SqlVirtualMachine/SqlVirtualMachines",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServerVirtualMachine),
		GetDescriber:         nil,
	},

	"Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.SqlVirtualMachine/SqlVirtualMachineGroups",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServerVirtualMachineGroups),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/tableServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/tableServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageTableService),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SynapseWorkspace),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces/bigdatapools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces/bigdatapools",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SynapseWorkspaceBigdataPools),
		GetDescriber:         nil,
	},

	"Microsoft.Synapse/workspaces/sqlpools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Synapse/workspaces/sqlpools",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Synapse%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SynapseWorkspaceSqlpools),
		GetDescriber:         nil,
	},

	"Microsoft.StreamAnalytics/streamingJobs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.StreamAnalytics/streamingJobs",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Stream%20Analytics%20job.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StreamAnalyticsJob),
		GetDescriber:         nil,
	},

	"Microsoft.CostManagement/CostBySubscription": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.CostManagement/CostBySubscription",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DailyCostBySubscription),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerService/managedClusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ContainerService/managedClusters",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/AKS%20Hybrid%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KubernetesCluster),
		GetDescriber:         nil,
	},

	"Microsoft.ContainerService/serviceVersions": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ContainerService/serviceVersions",
		Tags:                 map[string][]string{
            "category": {"Container"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KubernetesServiceVersion),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataFactory),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/servers/jobagents": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/servers/jobagents",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Elastic%20Job%20Agent.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServerJobAgents),
		GetDescriber:         nil,
	},

	"Microsoft.Security/autoProvisioningSettings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/autoProvisioningSettings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterAutoProvisioning),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/logProfiles": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Insights/logProfiles",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LogProfile),
		GetDescriber:         nil,
	},

	"Microsoft.DataBoxEdge/dataBoxEdgeDevices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataBoxEdge/dataBoxEdgeDevices",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Box%20Edge.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataboxEdgeDevice),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancer),
		GetDescriber:         nil,
	},

	"Microsoft.Network/azureFirewalls": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/azureFirewalls",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Firewall.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkAzureFirewall),
		GetDescriber:         nil,
	},

	"Microsoft.Management/locks": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Management/locks",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Lock.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ManagementLock),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets/networkInterfaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets/networkInterfaces",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Interface.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineScaleSetNetworkInterface),
		GetDescriber:         nil,
	},

	"Microsoft.Network/frontDoors": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/frontDoors",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Front%20Door.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.FrontDoor),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/policyAssignments": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Authorization/policyAssignments",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Assignment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PolicyAssignment),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/userEffectiveAccess": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Authorization/userEffectiveAccess",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Assignment.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.UserEffectiveAccess),
		GetDescriber:         nil,
	},

	"Microsoft.Search/searchServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Search/searchServices",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Search%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SearchService),
		GetDescriber:         nil,
	},

	"Microsoft.Security/settings": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/settings",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterSetting),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Recovery%20Services%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RecoveryServicesVault),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupJobs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupJobs",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RecoveryServicesBackupJobs),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupPolicies": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupPolicies",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RecoveryServicesBackupPolicies),
		GetDescriber:         nil,
	},

	"Microsoft.RecoveryServices/vaults/backupItems": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.RecoveryServices/vaults/backupItems",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RecoveryServicesBackupItem),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskEncryptionSets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskEncryptionSets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Disk%20Encryption%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskEncryptionSet),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/sqlDatabases": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/sqlDatabases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DocumentDBSQLDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.EventGrid/topics": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.EventGrid/topics",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Grid%20Topic.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.EventGridTopic),
		GetDescriber:         nil,
	},

	"Microsoft.EventHub/namespaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.EventHub/namespaces",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.EventhubNamespace),
		GetDescriber:         nil,
	},

	"Microsoft.EventHub/namespaces/eventHubs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.EventHub/namespaces/eventHubs",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.EventhubNamespaceEventhub),
		GetDescriber:         nil,
	},

	"Microsoft.MachineLearningServices/workspaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.MachineLearningServices/workspaces",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Machine%20Learning%20Studio%20Workspace%20(classic).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MachineLearningWorkspace),
		GetDescriber:         nil,
	},

	"Microsoft.Dashboard/grafana": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Dashboard/grafana",
		Tags:                 map[string][]string{
            "category": {"Managed Services"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Managed%20Grafana.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DashboardGrafana),
		GetDescriber:         nil,
	},

	"Microsoft.DesktopVirtualization/workspaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DesktopVirtualization/workspaces",
		Tags:                 map[string][]string{
            "category": {"End User"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Windows%20Virtual%20Desktop.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DesktopVirtualizationWorkspaces),
		GetDescriber:         nil,
	},

	"Microsoft.Network/trafficManagerProfiles": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/trafficManagerProfiles",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Traffic%20Manager%20profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.TrafficManagerProfile),
		GetDescriber:         nil,
	},

	"Microsoft.Network/dnsResolvers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/dnsResolvers",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DNS%20Private%20Resolver.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DNSResolvers),
		GetDescriber:         nil,
	},

	"Microsoft.CostManagement/CostByResourceType": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.CostManagement/CostByResourceType",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DailyCostByResourceType),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkInterfaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/networkInterfaces",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Interface.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkInterface),
		GetDescriber:         nil,
	},

	"Microsoft.Network/publicIPAddresses": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/publicIPAddresses",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Public%20IP%20Address.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PublicIPAddress),
		GetDescriber:         nil,
	},

	"Microsoft.HealthcareApis/services": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.HealthcareApis/services",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.HealthcareService),
		GetDescriber:         nil,
	},

	"Microsoft.ServiceBus/namespaces": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ServiceBus/namespaces",
		Tags:                 map[string][]string{
            "category": {"Intergration"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Service%20Bus.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ServicebusNamespace),
		GetDescriber:         nil,
	},

	"Microsoft.Web/sites": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/sites",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppServiceFunctionApp),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/availabilitySets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/availabilitySets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Availability%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeAvailabilitySet),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworks": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworks",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.VirtualNetwork),
		GetDescriber:         nil,
	},

	"Microsoft.Security/securityContacts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/securityContacts",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterContact),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskswriteops": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskswriteops",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskWriteOps),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskswriteopshourly": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskswriteopshourly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskWriteOpsHourly),
		GetDescriber:         nil,
	},

	"Microsoft.EventGrid/domains": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.EventGrid/domains",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Event%20Grid%20Domain.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.EventGridDomain),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/deletedVaults": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/deletedVaults",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DeletedVault),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/tableServices/tables": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/tableServices/tables",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageTable),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/snapshots": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/snapshots",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Managed%20Disk%20Snapshot.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeSnapshots),
		GetDescriber:         nil,
	},

	"Microsoft.Kusto/clusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Kusto/clusters",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_Grouped/Data/Azure%20Data%20Explorer%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KustoCluster),
		GetDescriber:         nil,
	},

	"Microsoft.StorageSync/storageSyncServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.StorageSync/storageSyncServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Sync%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageSync),
		GetDescriber:         nil,
	},

	"Microsoft.Security/locations/jitNetworkAccessPolicies": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/locations/jitNetworkAccessPolicies",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterJitNetworkAccessPolicy),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworks/subnets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworks/subnets",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Subnet.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.Subnet),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/backendAddressPools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/backendAddressPools",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Backend%20pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancerBackendAddressPool),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/loadBalancingRules": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/loadBalancingRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancerRule),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineCpuUtilizationDaily": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineCpuUtilizationDaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineCpuUtilizationDaily),
		GetDescriber:         nil,
	},

	"Microsoft.DataLakeStore/accounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataLakeStore/accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Lake.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataLakeStore),
		GetDescriber:         nil,
	},

	"Microsoft.StorageCache/caches": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.StorageCache/caches",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_Grouped/Data/HPC%20Cache.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.HpcCache),
		GetDescriber:         nil,
	},

	"Microsoft.Batch/batchAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Batch/batchAccounts",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Batch%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.BatchAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkSecurityGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/networkSecurityGroups",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Network%20Security%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkSecurityGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/roleDefinitions": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Authorization/roleDefinitions",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Role%20(Custom).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RoleDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Network/applicationSecurityGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/applicationSecurityGroups",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Security%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkApplicationSecurityGroups),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/roleAssignment": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Authorization/roleAssignment",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RoleAssignment),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DocumentDBMongoDatabase),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DocumentDBMongoCollection),
		GetDescriber:         nil,
	},

	"Microsoft.Network/networkWatchers/flowLogs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/networkWatchers/flowLogs",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkWatcherFlowLog),
		GetDescriber:         nil,
	},

	"microsoft.Sql/servers/elasticpools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "microsoft.Sql/servers/elasticpools",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Elastic%20Pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServerElasticPool),
		GetDescriber:         nil,
	},

	"Microsoft.Security/subAssessments": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/subAssessments",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterSubAssessment),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/disks": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/disks",
		Tags:                 map[string][]string{
            "category": {"Storage"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDisk),
		GetDescriber:         nil,
	},

	"Microsoft.Devices/ProvisioningServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Devices/ProvisioningServices",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/IoT%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.IOTHubDps),
		GetDescriber:         nil,
	},

	"Microsoft.HDInsight/clusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.HDInsight/clusters",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/HDInsight%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.HdInsightCluster),
		GetDescriber:         nil,
	},

	"Microsoft.ServiceFabric/clusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ServiceFabric/clusters",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Service%20Fabric%20Managed%20Cluster.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ServiceFabricCluster),
		GetDescriber:         nil,
	},

	"Microsoft.SignalRService/signalR": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.SignalRService/signalR",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SignalR.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SignalrService),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/blob": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/blob",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Blob.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageBlob),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageaccounts/blobservices/containers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageaccounts/blobservices/containers",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Container.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageContainer),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/blobServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/blobServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Blob.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageBlobService),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts/queueServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts/queueServices",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account%20Queue.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageQueue),
		GetDescriber:         nil,
	},

	"Microsoft.ApiManagement/service": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ApiManagement/service",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/API%20Management%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.APIManagement),
		GetDescriber:         nil,
	},

	"Microsoft.ApiManagement/backend": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.ApiManagement/backend",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/API%20Management%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.APIManagementBackend),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/disksreadops": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/disksreadops",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskReadOps),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineScaleSets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineScaleSets",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Scale%20Set.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineScaleSet),
		GetDescriber:         nil,
	},

	"Microsoft.DataFactory/factories/datasets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataFactory/factories/datasets",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Factory.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataFactoryDataset),
		GetDescriber:         nil,
	},

	"Microsoft.Authorization/policyDefinitions": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Authorization/policyDefinitions",
		Tags:                 map[string][]string{
            "category": {"Identify & Access"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Policy%20Definition.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PolicyDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/locations": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/locations",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.Location),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskAccesses": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskAccesses",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Disk%20Access.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskAccess),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMySQL/servers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforMySQL/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MySQL.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MysqlServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMySQL/flexibleservers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforMySQL/flexibleservers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Wordpress%20and%20MySQL%20Flexible%20server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MysqlFlexibleservers),
		GetDescriber:         nil,
	},

	"Microsoft.Cache/redisenterprise": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Cache/redisenterprise",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cache%20for%20Redis.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CacheRedisEnterprise),
		GetDescriber:         nil,
	},

	"Microsoft.DataLakeAnalytics/accounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DataLakeAnalytics/accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Data%20Lake%20Analytics.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DataLakeAnalyticsAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/activityLogAlerts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Insights/activityLogAlerts",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LogAlert),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineCpuUtilizationHourly": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineCpuUtilizationHourly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineCpuUtilizationHourly),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/outboundRules": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/outboundRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Backend%20Outbound%20Rule.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancerOutboundRule),
		GetDescriber:         nil,
	},

	"Microsoft.HybridCompute/machines": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.HybridCompute/machines",
		Tags:                 map[string][]string{
            "category": {"Compute"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.HybridComputeMachine),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/inboundNatRules": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/inboundNatRules",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Inbound%20NAT%20Rule.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancerNatRule),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/providers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/providers",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ResourceProvider),
		GetDescriber:         nil,
	},

	"Microsoft.Network/routeTables": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/routeTables",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Route%20Table.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.RouteTables),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/databaseAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/databaseAccounts",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CosmosdbAccount),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/restorableDatabaseAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/restorableDatabaseAccounts",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Cosmos%20DB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CosmosdbRestorableDatabaseAccount),
		GetDescriber:         nil,
	},

	"Microsoft.Network/applicationGateways": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/applicationGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ApplicationGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Security/automations": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Security/automations",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SecurityCenterAutomation),
		GetDescriber:         nil,
	},

	"Microsoft.Kubernetes/connectedClusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Kubernetes/connectedClusters",
		Tags:                 map[string][]string{
            "category": {"Container"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Kubernetes%20Cluster%20(Operator%20Nexus).svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.HybridKubernetesConnectedCluster),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/keys": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/keys",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVaultKey),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/certificates": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/certificates",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVaultCertificate),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/keys/Versions": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/keys/Versions",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVaultKey),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMariaDB/servers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforMariaDB/servers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MariaDB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MariadbServer),
		GetDescriber:         nil,
	},

	"Microsoft.DBforMariaDB/servers/databases": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DBforMariaDB/servers/databases",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Database%20for%20MariaDB.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MariadbDatabases),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/disksreadopsdaily": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/disksreadopsdaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskReadOpsDaily),
		GetDescriber:         nil,
	},

	"Microsoft.Web/plan": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Web/plan",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Service%20plan.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppServicePlan),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/disksreadopshourly": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/disksreadopshourly",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskReadOpsHourly),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/diskswriteopsdaily": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/diskswriteopsdaily",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeDiskWriteOpsDaily),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/tenants": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/tenants",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.Tenant),
		GetDescriber:         nil,
	},

	"Microsoft.Network/virtualNetworkGateways": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/virtualNetworkGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Network%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.VirtualNetworkGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Devices/iotHubs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Devices/iotHubs",
		Tags:                 map[string][]string{
            "category": {"IoT & Devices"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/IoT%20Hub.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.IOTHub),
		GetDescriber:         nil,
	},

	"Microsoft.Logic/workflows": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Logic/workflows",
		Tags:                 map[string][]string{
            "category": {"Integration"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LogicAppWorkflow),
		GetDescriber:         nil,
	},

	"Microsoft.Sql/flexibleServers": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Sql/flexibleServers",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SQL%20Server.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlServerFlexibleServer),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/links": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/links",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Resource%20Management%20Private%20Link.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ResourceLink),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions",
		Tags:                 map[string][]string{
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Subscription.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.Subscription),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/images": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/images",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine%20Image.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeImage),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachines": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachines",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Virtual%20Machine.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachine),
		GetDescriber:         nil,
	},

	"Microsoft.Network/natGateways": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/natGateways",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/NAT%20Gateway.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NatGateway),
		GetDescriber:         nil,
	},

	"Microsoft.Network/loadBalancers/probes": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/loadBalancers/probes",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Load%20Balancer%20Health%20Probe.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LoadBalancerProbe),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVault),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/managedHsms": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/managedHsms",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault%20HSM.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVaultManagedHardwareSecurityModule),
		GetDescriber:         nil,
	},

	"Microsoft.KeyVault/vaults/secrets": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.KeyVault/vaults/secrets",
		Tags:                 map[string][]string{
            "category": {"Security"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Key%20Vault%20Secret.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.KeyVaultSecret),
		GetDescriber:         nil,
	},

	"Microsoft.AppConfiguration/configurationStores": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.AppConfiguration/configurationStores",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/App%20Configuration.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.AppConfiguration),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/virtualMachineCpuUtilization": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/virtualMachineCpuUtilization",
		Tags:                 map[string][]string{
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeVirtualMachineCpuUtilization),
		GetDescriber:         nil,
	},

	"Microsoft.Storage/storageAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Storage/storageAccounts",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Storage%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.StorageAccount),
		GetDescriber:         nil,
	},

	"Microsoft.AppPlatform/Spring": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.AppPlatform/Spring",
		Tags:                 map[string][]string{
            "category": {"PaaS"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Spring%20Cloud.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SpringCloudService),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/galleries": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/galleries",
		Tags:                 map[string][]string{
            "category": {"General"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Compute%20Gallery.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeGallery),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/hostGroups": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/hostGroups",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Host%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeHostGroup),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/hostGroups/hosts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/hostGroups/hosts",
		Tags:                 map[string][]string{
            "category": {"Compute"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Host%20Group.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeHost),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/restorePointCollections": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/restorePointCollections",
		Tags:                 map[string][]string{
            "category": {"Backup"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeRestorePointCollection),
		GetDescriber:         nil,
	},

	"Microsoft.Compute/sshPublicKeys": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Compute/sshPublicKeys",
		Tags:                 map[string][]string{
            "category": {"Management & Governance"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/SSH%20key.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ComputeSSHPublicKey),
		GetDescriber:         nil,
	},

	"Microsoft.Cdn/profiles/endpoints": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Cdn/profiles/endpoints",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/CDN%20Profile.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.CdnEndpoint),
		GetDescriber:         nil,
	},

	"Microsoft.BotService/botServices": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.BotService/botServices",
		Tags:                 map[string][]string{
            "category": {"AI + ML"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Bot%20Service.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.BotServiceBot),
		GetDescriber:         nil,
	},

	"Microsoft.DocumentDB/cassandraClusters": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DocumentDB/cassandraClusters",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20Managed%20Instance%20for%20Apache%20Cassandra.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DocumentDBCassandraCluster),
		GetDescriber:         nil,
	},

	"Microsoft.Network/ddosProtectionPlans": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Network/ddosProtectionPlans",
		Tags:                 map[string][]string{
            "category": {"Networking"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DDoS%20Protection%20Plan.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetworkDDoSProtectionPlan),
		GetDescriber:         nil,
	},

	"microsoft.Sql/instancePools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "microsoft.Sql/instancePools",
		Tags:                 map[string][]string{
            "category": {"Database"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Instance%20Pool.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.SqlInstancePool),
		GetDescriber:         nil,
	},

	"microsoft.NetApp/netAppAccounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "microsoft.NetApp/netAppAccounts",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20NetApp%20Files.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetAppAccount),
		GetDescriber:         nil,
	},

	"Microsoft.NetApp/netAppAccounts/capacityPools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.NetApp/netAppAccounts/capacityPools",
		Tags:                 map[string][]string{
            "category": {"Storage"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Azure%20NetApp%20Files.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.NetAppCapacityPool),
		GetDescriber:         nil,
	},

	"Microsoft.DesktopVirtualization/hostpools": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.DesktopVirtualization/hostpools",
		Tags:                 map[string][]string{
            "category": {"End User"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Windows%20Virtual%20Desktop.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DesktopVirtualizationHostPool),
		GetDescriber:         nil,
	},

	"Microsoft.Devtestlab/labs": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Devtestlab/labs",
		Tags:                 map[string][]string{
            "category": {"DevOps + Testing"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/DevTest%20Lab.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.DevTestLabLab),
		GetDescriber:         nil,
	},

	"Microsoft.Purview/Accounts": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Purview/Accounts",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Purview%20Account.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PurviewAccount),
		GetDescriber:         nil,
	},

	"Microsoft.PowerBIDedicated/capacities": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.PowerBIDedicated/capacities",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Power%20BI%20Embedded.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.PowerBIDedicatedCapacity),
		GetDescriber:         nil,
	},

	"Microsoft.Insights/components": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Insights/components",
		Tags:                 map[string][]string{
            "category": {"Data and Analytics"},
            "logo_uri": {"https://raw.githubusercontent.com/opengovernance-io/Azure-Design/master/SVG_Azure_All/Application%20Insights.svg"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.ApplicationInsights),
		GetDescriber:         nil,
	},

	"Microsoft.Lighthouse/definition": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Lighthouse/definition",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LighthouseDefinition),
		GetDescriber:         nil,
	},

	"Microsoft.Lighthouse/assignment": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Lighthouse/assignment",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.LighthouseAssignments),
		GetDescriber:         nil,
	},

	"Microsoft.Maintenance/maintenanceConfigurations": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Maintenance/maintenanceConfigurations",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MaintenanceConfiguration),
		GetDescriber:         nil,
	},

	"Microsoft.Monitor/logProfiles": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Monitor/logProfiles",
		Tags:                 map[string][]string{
            "category": {},
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.MonitorLogProfiles),
		GetDescriber:         nil,
	},

	"Microsoft.Resources/subscriptions/resources": {
		IntegrationType:      configs.IntegrationName,
		ResourceName:         "Microsoft.Resources/subscriptions/resources",
		Tags:                 map[string][]string{
            "logo_uri": {},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        DescribeBySubscription(describer.Resources),
		GetDescriber:         nil,
	},
}
