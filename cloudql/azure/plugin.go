package azure

import (
	"context"

	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-azure"

// Plugin creates this (azure) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound"}),
			},
		},
		// Default ignore config for the plugin
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrorPluginDefault(),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		TableMap: map[string]*plugin.Table{
			"azure_app_containerapps":               tableAzureAppContainerApps(ctx),
			"azure_app_managedenvironments":         tableAzureAppManagedEnvironments(ctx),
			"azure_blueprint_blueprints":            tableAzureBlueprintBlueprints(ctx),
			"azure_botservice_bot":                  tableAzureBotServiceBot(ctx),
			"azure_cdn_endpoint":                    tableAzureCdnEndpoint(ctx),
			"azure_cdn_profiles":                    tableAzureCdnProfiles(ctx),
			"azure_compute_cloudservices":           tableAzureComputeCloudServices(ctx),
			"azure_datamigration_services":          tableAzureDataMigrationServices(ctx),
			"azure_dashboard_grafana":               tableAzureDashboardGrafana(ctx),
			"azure_desktopvirtualization_workspace": tableAzureDesktopVirtualizationWorkspace(ctx),
			"azure_network_dnsresolver":             tableAzureNetworkDNSResolver(ctx),
			"azure_trafficmanager_profile":          tableAzureTrafficManagerProfile(ctx),
			"azure_dataprotection_backuppolicies":   tableAzureDataProtectionBackupPolicies(ctx),
			"azure_dataprotection_backupvaults":     tableAzureDataProtectionBackupVaults(ctx),
			"azure_logic_integrationaccounts":       tableAzureLogicIntegrationAccounts(ctx),
			"azure_network_connections":             tableAzureNetworkConnections(ctx),
			"azure_network_ddos_protection_plan":    tableAzureNetworkDDoSProtectionPlan(ctx),
			"azure_network_localnetworkgateways":    tableAzureNetworkLocalNetworkGateways(ctx),
			"azure_network_privatelinkservices":     tableAzureNetworkPrivateLinkServices(ctx),
			"azure_network_publicipprefixes":        tableAzureNetworkPublicIPPrefixes(ctx),
			"azure_network_virtualhubs":             tableAzureNetworkVirtualHubs(ctx),
			"azure_network_virtualwans":             tableAzureNetworkVirtualWans(ctx),
			"azure_network_vpnconnections":          tableAzureNetworkVPNConnections(ctx),
			"azure_network_vpngateways":             tableAzureNetworkVPNGateways(ctx),
			"azure_network_vpnsites":                tableAzureNetworkVPNSites(ctx),
			"azure_streamanalytics_cluster":         tableAzureStreamAnalyticsCluster(ctx),
			//"azure_timeseriesinsights_environments":                     tableAzureTimeSeriesInsightsEnvironments(ctx),
			"azure_virtualmachineimages_imagetemplates": tableAzureVirtualMachineImagesImageTemplates(ctx),
			"azure_web_serverfarms":                     tableAzureWebServerFarms(ctx),
			"azure_api_management":                      tableAzureAPIManagement(ctx),
			"azure_api_management_backend":              tableAzureAPIManagementBackend(ctx),
			"azure_app_configuration":                   tableAzureAppConfiguration(ctx),
			"azure_app_service_environment":             tableAzureAppServiceEnvironment(ctx),
			"azure_app_service_function_app":            tableAzureAppServiceFunctionApp(ctx),
			"azure_app_service_plan":                    tableAzureAppServicePlan(ctx),
			"azure_app_service_web_app":                 tableAzureAppServiceWebApp(ctx),
			"azure_app_service_web_app_slot":            tableAzureAppServiceWebAppSlot(ctx),
			"azure_application_gateway":                 tableAzureApplicationGateway(ctx),
			"azure_application_insight":                 tableAzureApplicationInsight(ctx),
			"azure_application_security_group":          tableAzureApplicationSecurityGroup(ctx),
			"azure_automation_account":                  tableAzureApAutomationAccount(ctx),
			"azure_automation_variable":                 tableAzureApAutomationVariable(ctx),
			"azure_bastion_host":                        tableAzureBastionHost(ctx),
			"azure_batch_account":                       tableAzureBatchAccount(ctx),
			"azure_cognitive_account":                   tableAzureCognitiveAccount(ctx),
			"azure_compute_availability_set":            tableAzureComputeAvailabilitySet(ctx),
			"azure_compute_disk":                        tableAzureComputeDisk(ctx),
			"azure_compute_disk_access":                 tableAzureComputeDiskAccess(ctx),
			"azure_compute_disk_encryption_set":         tableAzureComputeDiskEncryptionSet(ctx),
			"azure_compute_host_group":                  tableAzureComputeHostGroup(ctx),
			"azure_compute_host":                        tableAzureComputeHost(ctx),
			"azure_compute_image":                       tableAzureComputeImage(ctx),
			"azure_compute_image_gallery":               tableAzureComputeImageGallery(ctx),
			//"azure_compute_resource_sku":                                tableAzureResourceSku(ctx),
			"azure_compute_restore_point_collection":                    tableAzureComputeRestorePointCollection(ctx),
			"azure_compute_snapshot":                                    tableAzureComputeSnapshot(ctx),
			"azure_compute_ssh_key":                                     tableAzureComputeSshKey(ctx),
			"azure_compute_virtual_machine":                             tableAzureComputeVirtualMachine(ctx),
			"azure_compute_virtual_machine_scale_set":                   tableAzureComputeVirtualMachineScaleSet(ctx),
			"azure_compute_virtual_machine_scale_set_network_interface": tableAzureComputeVirtualMachineScaleSetNetworkInterface(ctx),
			"azure_compute_virtual_machine_scale_set_vm":                tableAzureComputeVirtualMachineScaleSetVm(ctx),
			"azure_container_group":                                     tableAzureContainerGroup(ctx),
			"azure_container_registry":                                  tableAzureContainerRegistry(ctx),
			"azure_cosmosdb_account":                                    tableAzureCosmosDBAccount(ctx),
			"azure_cosmosdb_mongo_collection":                           tableAzureCosmosDBMongoCollection(ctx),
			"azure_cosmosdb_mongo_database":                             tableAzureCosmosDBMongoDatabase(ctx),
			"azure_cosmosdb_restorable_database_account":                tableAzureCosmosDBRestorableDatabaseAccount(ctx),
			"azure_cosmosdb_sql_database":                               tableAzureCosmosDBSQLDatabase(ctx),
			"azure_cosmosdb_cassandra_cluster":                          tableAzureCosmosdbCassandraCluster(ctx),
			"azure_costmanagement_costbyresourcetype":                   tableAzureCostManagementCostByResourceType(ctx),
			"azure_costmanagement_costbysubscription":                   tableAzureCostManagementCostBySubscription(ctx),
			"azure_data_factory":                                        tableAzureDataFactory(ctx),
			"azure_data_factory_dataset":                                tableAzureDataFactoryDataset(ctx),
			"azure_data_factory_pipeline":                               tableAzureDataFactoryPipeline(ctx),
			"azure_data_lake_analytics_account":                         tableAzureDataLakeAnalyticsAccount(ctx),
			//"azure_data_lake_store":                                     tableAzureDataLakeStore(ctx),
			"azure_databox_edge_device":                        tableAzureDataBoxEdgeDevice(ctx),
			"azure_diagnostic_setting":                         tableAzureDiagnosticSetting(ctx),
			"azure_autoscale_setting":                          tableAzureAutoscaleSetting(ctx),
			"azure_dns_zone":                                   tableAzureDNSZone(ctx),
			"azure_eventgrid_domain":                           tableAzureEventGridDomain(ctx),
			"azure_eventgrid_topic":                            tableAzureEventGridTopic(ctx),
			"azure_eventhub_namespace":                         tableAzureEventHubNamespace(ctx),
			"azure_express_route_circuit":                      tableAzureExpressRouteCircuit(ctx),
			"azure_firewall":                                   tableAzureFirewall(ctx),
			"azure_firewall_policy":                            tableAzureFirewallPolicy(ctx),
			"azure_frontdoor":                                  tableAzureFrontDoor(ctx),
			"azure_hdinsight_cluster":                          tableAzureHDInsightCluster(ctx),
			"azure_healthcare_service":                         tableAzureHealthcareService(ctx),
			"azure_hpc_cache":                                  tableAzureHPCCache(ctx),
			"azure_hybrid_compute_machine":                     tableAzureHybridComputeMachine(ctx),
			"azure_hybrid_kubernetes_connected_cluster":        tableAzureHybridKubernetesConnectedCluster(ctx),
			"azure_iothub":                                     tableAzureIotHub(ctx),
			"azure_iothub_dps":                                 tableAzureIotHubDps(ctx),
			"azure_key_vault":                                  tableAzureKeyVault(ctx),
			"azure_key_vault_deleted_vault":                    tableAzureKeyVaultDeletedVault(ctx),
			"azure_key_vault_key":                              tableAzureKeyVaultKey(ctx),
			"azure_key_vault_certificate":                      tableAzureKeyVaultCertificate(ctx),
			"azure_key_vault_key_version":                      tableAzureKeyVaultKeyVersion(ctx),
			"azure_key_vault_managed_hardware_security_module": tableAzureKeyVaultManagedHardwareSecurityModule(ctx),
			"azure_key_vault_secret":                           tableAzureKeyVaultSecret(ctx),
			"azure_kubernetes_cluster":                         tableAzureKubernetesCluster(ctx),
			"azure_kubernetes_service_version":                 tableAzureAKSVersion(ctx),
			"azure_kusto_cluster":                              tableAzureKustoCluster(ctx),
			"azure_lb":                                         tableAzureLoadBalancer(ctx),
			"azure_lb_backend_address_pool":                    tableAzureLoadBalancerBackendAddressPool(ctx),
			"azure_lb_nat_rule":                                tableAzureLoadBalancerNatRule(ctx),
			"azure_lb_outbound_rule":                           tableAzureLoadBalancerOutboundRule(ctx),
			"azure_lb_probe":                                   tableAzureLoadBalancerProbe(ctx),
			"azure_lb_rule":                                    tableAzureLoadBalancerRule(ctx),
			"azure_location":                                   tableAzureLocation(ctx),
			"azure_log_alert":                                  tableAzureLogAlert(ctx),
			"azure_log_profile":                                tableAzureLogProfile(ctx),
			"azure_logic_app_workflow":                         tableAzureLogicAppWorkflow(ctx),
			"azure_machine_learning_workspace":                 tableAzureMachineLearningWorkspace(ctx),
			"azure_management_group":                           tableAzureManagementGroup(ctx),
			"azure_management_lock":                            tableAzureManagementLock(ctx),
			"azure_mariadb_server":                             tableAzureMariaDBServer(ctx),
			"azure_mssql_elasticpool":                          tableAzureMSSQLElasticPool(ctx),
			"azure_mssql_managed_instance":                     tableAzureMSSQLManagedInstance(ctx),
			"azure_mssql_virtual_machine":                      tableAzureMSSQLVirtualMachine(ctx),
			"azure_mysql_flexible_server":                      tableAzureMySQLFlexibleServer(ctx),
			"azure_mysql_server":                               tableAzureMySQLServer(ctx),
			"azure_nat_gateway":                                tableAzureNatGateway(ctx),
			"azure_netapp_account":                             tableAzureNetAppAccount(ctx),
			"azure_netapp_capacity_pool":                       tableAzureNetAppCapacityPool(ctx),
			"azure_network_interface":                          tableAzureNetworkInterface(ctx),
			"azure_network_security_group":                     tableAzureNetworkSecurityGroup(ctx),
			"azure_network_watcher":                            tableAzureNetworkWatcher(ctx),
			"azure_network_watcher_flow_log":                   tableAzureNetworkWatcherFlowLog(ctx),
			"azure_policy_assignment":                          tableAzurePolicyAssignment(ctx),
			"azure_policy_definition":                          tableAzurePolicyDefinition(ctx),
			"azure_postgresql_server":                          tableAzurePostgreSqlServer(ctx),
			"azure_powerbidedicated_capacity":                  tableAzurePowerBIDedicatedCapacity(ctx),
			"azure_private_dns_zone":                           tableAzurePrivateDNSZone(ctx),
			"azure_provider":                                   tableAzureProvider(ctx),
			"azure_public_ip":                                  tableAzurePublicIP(ctx),
			"azure_purview_account":                            tableAzurePurviewAccount(ctx),
			"azure_recovery_services_vault":                    tableAzureRecoveryServicesVault(ctx),
			"azure_recovery_services_backup_job":               tableAzureRecoveryServicesBackupJob(ctx),
			"azure_recovery_services_backup_policy":            tableAzureRecoveryServicesBackupPolicy(ctx),
			"azure_recovery_services_backup_item":              tableAzureRecoveryServicesBackupItem(ctx),
			"azure_redis_cache":                                tableAzureRedisCache(ctx),
			"azure_resource_group":                             tableAzureResourceGroup(ctx),
			"azure_resource_link":                              tableAzureResourceLink(ctx),
			"azure_role_assignment":                            tableAzureIamRoleAssignment(ctx),
			"azure_role_definition":                            tableAzureIamRoleDefinition(ctx),
			"azure_route_table":                                tableAzureRouteTable(ctx),
			"azure_search_service":                             tableAzureSearchService(ctx),
			"azure_security_center_automation":                 tableAzureSecurityCenterAutomation(ctx),
			"azure_security_center_auto_provisioning":          tableAzureSecurityCenterAutoProvisioning(ctx),
			"azure_security_center_contact":                    tableAzureSecurityCenterContact(ctx),
			"azure_security_center_jit_network_access_policy":  tableAzureSecurityCenterJITNetworkAccessPolicy(ctx),
			"azure_security_center_setting":                    tableAzureSecurityCenterSetting(ctx),
			"azure_security_center_sub_assessment":             tableAzureSecurityCenterSubAssessment(ctx),
			//"azure_security_center_subscription_pricing":                tableAzureSecurityCenterPricing(ctx),
			"azure_service_fabric_cluster":           tableAzureServiceFabricCluster(ctx),
			"azure_servicebus_namespace":             tableAzureServiceBusNamespace(ctx),
			"azure_signalr_service":                  tableAzureSignalRService(ctx),
			"azure_spring_cloud_service":             tableAzureSpringCloudService(ctx),
			"azure_sql_database":                     tableAzureSqlDatabase(ctx),
			"azure_sql_instance_pool":                tableAzureSqlInstancePool(ctx),
			"azure_sql_server":                       tableAzureSQLServer(ctx),
			"azure_storage_account":                  tableAzureStorageAccount(ctx),
			"azure_storage_blob":                     tableAzureStorageBlob(ctx),
			"azure_storage_blob_service":             tableAzureStorageBlobService(ctx),
			"azure_storage_container":                tableAzureStorageContainer(ctx),
			"azure_storage_queue":                    tableAzureStorageQueue(ctx),
			"azure_storage_share_file":               tableAzureStorageShareFile(ctx),
			"azure_storage_sync":                     tableAzureStorageSync(ctx),
			"azure_storage_table":                    tableAzureStorageTable(ctx),
			"azure_storage_table_service":            tableAzureStorageTableService(ctx),
			"azure_stream_analytics_job":             tableAzureStreamAnalyticsJob(ctx),
			"azure_subnet":                           tableAzureSubnet(ctx),
			"azure_subscription":                     tableAzureSubscription(ctx),
			"azure_synapse_workspace":                tableAzureSynapseWorkspace(ctx),
			"azure_tenant":                           tableAzureTenant(ctx),
			"azure_virtual_network":                  tableAzureVirtualNetwork(ctx),
			"azure_virtual_network_gateway":          tableAzureVirtualNetworkGateway(ctx),
			"azure_network_privateendpoints":         tableAzureNetworkPrivateEndpoints(ctx),
			"azure_eventhub_namespaceeventhubs":      tableAzureEventhubNamespaceEventhubs(ctx),
			"azure_mariadb_databases":                tableAzureMariaDBDatabases(ctx),
			"azure_analysisservices_servers":         tableAzureAnalysisServicesServers(ctx),
			"azure_cache_redisenterprise":            tableAzureCacheRedisEnterprise(ctx),
			"azure_databricks_workspaces":            tableAzureDatabricksWorkspaces(ctx),
			"azure_dbformysql_flexibleservers":       tableAzureDBforMySQLFlexibleServers(ctx),
			"azure_postgresql_flexible_server":       tableAzureDBforPostgreSQLFlexibleServers(ctx),
			"azure_desktop_virtualization_host_pool": tableAzureDesktopVirtualizationHostPool(ctx),
			"azure_devtestlab_lab":                   tableAzureDevTestLabLab(ctx),
			"azure_sql_managedinstancesdatabases":    tableAzureSqlManagedInstancesDatabases(ctx),
			"azure_sql_serversjobagents":             tableAzureSqlServersJobAgents(ctx),
			"azure_sql_virtualclusters":              tableAzureSqlVirtualClusters(ctx),
			"azure_sql_virtualmachinegroups":         tableAzureSqlVirtualMachineGroups(ctx),
			"azure_synapse_workspacesbigdatapools":   tableAzureSynapseWorkspacesBigDataPools(ctx),
			"azure_synapse_workspacessqlpools":       tableAzureSynapseWorkspacesSqlPools(ctx),
			"azure_user_effective_access":            tableAzureUserEffectiveAccess(ctx),
			"azure_lighthouse_definition":            tableAzureLighthouseDefinition(ctx),
			"azure_lighthouse_assignment":            tableAzureLighthouseAssignment(ctx),
			"azure_maintenance_configuration":        tableAzureMaintenanceConfiguration(ctx),
			"azure_monitor_log_profile":              tableAzureMonitorLogProfile(ctx),
			"azure_private_endpoint":                 tableAzurePrivateEndpoint(ctx),
			"azure_data_protection_backup_job":       tableAzureDataProtectionBackupJob(ctx),
			"azure_resource":                         tableAzureResourceResource(ctx),
			"azure_alert_management":                 tableAzureAlertMangement(ctx),
			"azure_log_analytics_workspace":          tableAzureLogAnalyticsWorkspace(ctx),
			"azure_operationalinsights_workspaces":   tableAzureOperationalInsightsWorkspaces(ctx),
		},
	}

	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_integration_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_integration_id", "platform_resource_id"})...)
			}
		}
	}

	return p
}
