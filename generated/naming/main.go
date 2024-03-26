// naming
package naming

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"naming.Naming",
		reflect.TypeOf((*Naming)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "analysisServicesServerOutput", GoGetter: "AnalysisServicesServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "apiManagementOutput", GoGetter: "ApiManagementOutput"},
			_jsii_.MemberProperty{JsiiProperty: "appConfigurationOutput", GoGetter: "AppConfigurationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationGatewayOutput", GoGetter: "ApplicationGatewayOutput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationInsightsOutput", GoGetter: "ApplicationInsightsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationSecurityGroupOutput", GoGetter: "ApplicationSecurityGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "appServiceOutput", GoGetter: "AppServiceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "appServicePlanOutput", GoGetter: "AppServicePlanOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationAccountOutput", GoGetter: "AutomationAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationCertificateOutput", GoGetter: "AutomationCertificateOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationCredentialOutput", GoGetter: "AutomationCredentialOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationRunbookOutput", GoGetter: "AutomationRunbookOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationScheduleOutput", GoGetter: "AutomationScheduleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "automationVariableOutput", GoGetter: "AutomationVariableOutput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilitySetOutput", GoGetter: "AvailabilitySetOutput"},
			_jsii_.MemberProperty{JsiiProperty: "bastionHostOutput", GoGetter: "BastionHostOutput"},
			_jsii_.MemberProperty{JsiiProperty: "batchAccountOutput", GoGetter: "BatchAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "batchApplicationOutput", GoGetter: "BatchApplicationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "batchCertificateOutput", GoGetter: "BatchCertificateOutput"},
			_jsii_.MemberProperty{JsiiProperty: "batchPoolOutput", GoGetter: "BatchPoolOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botChannelDirectlineOutput", GoGetter: "BotChannelDirectlineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botChannelEmailOutput", GoGetter: "BotChannelEmailOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botChannelMsTeamsOutput", GoGetter: "BotChannelMsTeamsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botChannelSlackOutput", GoGetter: "BotChannelSlackOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botChannelsRegistrationOutput", GoGetter: "BotChannelsRegistrationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botConnectionOutput", GoGetter: "BotConnectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "botWebAppOutput", GoGetter: "BotWebAppOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cdnEndpointOutput", GoGetter: "CdnEndpointOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cdnProfileOutput", GoGetter: "CdnProfileOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cognitiveAccountOutput", GoGetter: "CognitiveAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerGroupOutput", GoGetter: "ContainerGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "containerRegistryOutput", GoGetter: "ContainerRegistryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "containerRegistryWebhookOutput", GoGetter: "ContainerRegistryWebhookOutput"},
			_jsii_.MemberProperty{JsiiProperty: "cosmosdbAccountOutput", GoGetter: "CosmosdbAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "customProviderOutput", GoGetter: "CustomProviderOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardOutput", GoGetter: "DashboardOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseMigrationProjectOutput", GoGetter: "DatabaseMigrationProjectOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databaseMigrationServiceOutput", GoGetter: "DatabaseMigrationServiceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksClusterOutput", GoGetter: "DatabricksClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksHighConcurrencyClusterOutput", GoGetter: "DatabricksHighConcurrencyClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksStandardClusterOutput", GoGetter: "DatabricksStandardClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "databricksWorkspaceOutput", GoGetter: "DatabricksWorkspaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryDatasetMysqlOutput", GoGetter: "DataFactoryDatasetMysqlOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryDatasetPostgresqlOutput", GoGetter: "DataFactoryDatasetPostgresqlOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryDatasetSqlServerTableOutput", GoGetter: "DataFactoryDatasetSqlServerTableOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryIntegrationRuntimeManagedOutput", GoGetter: "DataFactoryIntegrationRuntimeManagedOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryLinkedServiceDataLakeStorageGen2Output", GoGetter: "DataFactoryLinkedServiceDataLakeStorageGen2Output"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryLinkedServiceKeyVaultOutput", GoGetter: "DataFactoryLinkedServiceKeyVaultOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryLinkedServiceMysqlOutput", GoGetter: "DataFactoryLinkedServiceMysqlOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryLinkedServicePostgresqlOutput", GoGetter: "DataFactoryLinkedServicePostgresqlOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryLinkedServiceSqlServerOutput", GoGetter: "DataFactoryLinkedServiceSqlServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryOutput", GoGetter: "DataFactoryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryPipelineOutput", GoGetter: "DataFactoryPipelineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataFactoryTriggerScheduleOutput", GoGetter: "DataFactoryTriggerScheduleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataLakeAnalyticsAccountOutput", GoGetter: "DataLakeAnalyticsAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataLakeAnalyticsFirewallRuleOutput", GoGetter: "DataLakeAnalyticsFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataLakeStoreFirewallRuleOutput", GoGetter: "DataLakeStoreFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dataLakeStoreOutput", GoGetter: "DataLakeStoreOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "devTestLabOutput", GoGetter: "DevTestLabOutput"},
			_jsii_.MemberProperty{JsiiProperty: "devTestLinuxVirtualMachineOutput", GoGetter: "DevTestLinuxVirtualMachineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "devTestWindowsVirtualMachineOutput", GoGetter: "DevTestWindowsVirtualMachineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionSetOutput", GoGetter: "DiskEncryptionSetOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsAaaaRecordOutput", GoGetter: "DnsAaaaRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsARecordOutput", GoGetter: "DnsARecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsCaaRecordOutput", GoGetter: "DnsCaaRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsCnameRecordOutput", GoGetter: "DnsCnameRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsMxRecordOutput", GoGetter: "DnsMxRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsNsRecordOutput", GoGetter: "DnsNsRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsPtrRecordOutput", GoGetter: "DnsPtrRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsTxtRecordOutput", GoGetter: "DnsTxtRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsZoneOutput", GoGetter: "DnsZoneOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventgridDomainOutput", GoGetter: "EventgridDomainOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventgridDomainTopicOutput", GoGetter: "EventgridDomainTopicOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventgridEventSubscriptionOutput", GoGetter: "EventgridEventSubscriptionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventgridTopicOutput", GoGetter: "EventgridTopicOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubAuthorizationRuleOutput", GoGetter: "EventhubAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubConsumerGroupOutput", GoGetter: "EventhubConsumerGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubNamespaceAuthorizationRuleOutput", GoGetter: "EventhubNamespaceAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubNamespaceDisasterRecoveryConfigOutput", GoGetter: "EventhubNamespaceDisasterRecoveryConfigOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubNamespaceOutput", GoGetter: "EventhubNamespaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "eventhubOutput", GoGetter: "EventhubOutput"},
			_jsii_.MemberProperty{JsiiProperty: "expressRouteCircuitOutput", GoGetter: "ExpressRouteCircuitOutput"},
			_jsii_.MemberProperty{JsiiProperty: "expressRouteGatewayOutput", GoGetter: "ExpressRouteGatewayOutput"},
			_jsii_.MemberProperty{JsiiProperty: "firewallApplicationRuleCollectionOutput", GoGetter: "FirewallApplicationRuleCollectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "firewallIpConfigurationOutput", GoGetter: "FirewallIpConfigurationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "firewallNatRuleCollectionOutput", GoGetter: "FirewallNatRuleCollectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "firewallNetworkRuleCollectionOutput", GoGetter: "FirewallNetworkRuleCollectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "firewallOutput", GoGetter: "FirewallOutput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "frontdoorFirewallPolicyOutput", GoGetter: "FrontdoorFirewallPolicyOutput"},
			_jsii_.MemberProperty{JsiiProperty: "frontdoorOutput", GoGetter: "FrontdoorOutput"},
			_jsii_.MemberProperty{JsiiProperty: "functionAppOutput", GoGetter: "FunctionAppOutput"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightHadoopClusterOutput", GoGetter: "HdinsightHadoopClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightHbaseClusterOutput", GoGetter: "HdinsightHbaseClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightInteractiveQueryClusterOutput", GoGetter: "HdinsightInteractiveQueryClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightKafkaClusterOutput", GoGetter: "HdinsightKafkaClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightMlServicesClusterOutput", GoGetter: "HdinsightMlServicesClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightRserverClusterOutput", GoGetter: "HdinsightRserverClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightSparkClusterOutput", GoGetter: "HdinsightSparkClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "hdinsightStormClusterOutput", GoGetter: "HdinsightStormClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "imageOutput", GoGetter: "ImageOutput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "iotcentralApplicationOutput", GoGetter: "IotcentralApplicationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "iothubConsumerGroupOutput", GoGetter: "IothubConsumerGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "iothubDpsCertificateOutput", GoGetter: "IothubDpsCertificateOutput"},
			_jsii_.MemberProperty{JsiiProperty: "iothubDpsOutput", GoGetter: "IothubDpsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "iothubOutput", GoGetter: "IothubOutput"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultCertificateOutput", GoGetter: "KeyVaultCertificateOutput"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultKeyOutput", GoGetter: "KeyVaultKeyOutput"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultOutput", GoGetter: "KeyVaultOutput"},
			_jsii_.MemberProperty{JsiiProperty: "keyVaultSecretOutput", GoGetter: "KeyVaultSecretOutput"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesClusterOutput", GoGetter: "KubernetesClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "kustoClusterOutput", GoGetter: "KustoClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "kustoDatabaseOutput", GoGetter: "KustoDatabaseOutput"},
			_jsii_.MemberProperty{JsiiProperty: "kustoEventhubDataConnectionOutput", GoGetter: "KustoEventhubDataConnectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "lbNatRuleOutput", GoGetter: "LbNatRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "lbOutput", GoGetter: "LbOutput"},
			_jsii_.MemberProperty{JsiiProperty: "linuxVirtualMachineOutput", GoGetter: "LinuxVirtualMachineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "linuxVirtualMachineScaleSetOutput", GoGetter: "LinuxVirtualMachineScaleSetOutput"},
			_jsii_.MemberProperty{JsiiProperty: "localNetworkGatewayOutput", GoGetter: "LocalNetworkGatewayOutput"},
			_jsii_.MemberProperty{JsiiProperty: "logAnalyticsWorkspaceOutput", GoGetter: "LogAnalyticsWorkspaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "machineLearningWorkspaceOutput", GoGetter: "MachineLearningWorkspaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "managedDiskOutput", GoGetter: "ManagedDiskOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mapsAccountOutput", GoGetter: "MapsAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbDatabaseOutput", GoGetter: "MariadbDatabaseOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbFirewallRuleOutput", GoGetter: "MariadbFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbServerOutput", GoGetter: "MariadbServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mariadbVirtualNetworkRuleOutput", GoGetter: "MariadbVirtualNetworkRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mssqlDatabaseOutput", GoGetter: "MssqlDatabaseOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mssqlElasticpoolOutput", GoGetter: "MssqlElasticpoolOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mssqlServerOutput", GoGetter: "MssqlServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabaseOutput", GoGetter: "MysqlDatabaseOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlFirewallRuleOutput", GoGetter: "MysqlFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlServerOutput", GoGetter: "MysqlServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlVirtualNetworkRuleOutput", GoGetter: "MysqlVirtualNetworkRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkDdosProtectionPlanOutput", GoGetter: "NetworkDdosProtectionPlanOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceOutput", GoGetter: "NetworkInterfaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupOutput", GoGetter: "NetworkSecurityGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupRuleOutput", GoGetter: "NetworkSecurityGroupRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityRuleOutput", GoGetter: "NetworkSecurityRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "networkWatcherOutput", GoGetter: "NetworkWatcherOutput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationHubAuthorizationRuleOutput", GoGetter: "NotificationHubAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationHubNamespaceOutput", GoGetter: "NotificationHubNamespaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationHubOutput", GoGetter: "NotificationHubOutput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pointToSiteVpnGatewayOutput", GoGetter: "PointToSiteVpnGatewayOutput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlDatabaseOutput", GoGetter: "PostgresqlDatabaseOutput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlFirewallRuleOutput", GoGetter: "PostgresqlFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlServerOutput", GoGetter: "PostgresqlServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlVirtualNetworkRuleOutput", GoGetter: "PostgresqlVirtualNetworkRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "powerbiEmbeddedOutput", GoGetter: "PowerbiEmbeddedOutput"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsAaaaRecordOutput", GoGetter: "PrivateDnsAaaaRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsARecordOutput", GoGetter: "PrivateDnsARecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsCnameRecordOutput", GoGetter: "PrivateDnsCnameRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsMxRecordOutput", GoGetter: "PrivateDnsMxRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsPtrRecordOutput", GoGetter: "PrivateDnsPtrRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsSrvRecordOutput", GoGetter: "PrivateDnsSrvRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsTxtRecordOutput", GoGetter: "PrivateDnsTxtRecordOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsZoneGroupOutput", GoGetter: "PrivateDnsZoneGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsZoneOutput", GoGetter: "PrivateDnsZoneOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpointOutput", GoGetter: "PrivateEndpointOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateLinkServiceOutput", GoGetter: "PrivateLinkServiceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "privateServiceConnectionOutput", GoGetter: "PrivateServiceConnectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "proximityPlacementGroupOutput", GoGetter: "ProximityPlacementGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpOutput", GoGetter: "PublicIpOutput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpPrefixOutput", GoGetter: "PublicIpPrefixOutput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "redisCacheOutput", GoGetter: "RedisCacheOutput"},
			_jsii_.MemberProperty{JsiiProperty: "redisFirewallRuleOutput", GoGetter: "RedisFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "relayHybridConnectionOutput", GoGetter: "RelayHybridConnectionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "relayNamespaceOutput", GoGetter: "RelayNamespaceOutput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupOutput", GoGetter: "ResourceGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "roleAssignmentOutput", GoGetter: "RoleAssignmentOutput"},
			_jsii_.MemberProperty{JsiiProperty: "roleDefinitionOutput", GoGetter: "RoleDefinitionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "routeOutput", GoGetter: "RouteOutput"},
			_jsii_.MemberProperty{JsiiProperty: "routeTableOutput", GoGetter: "RouteTableOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusNamespaceAuthorizationRuleOutput", GoGetter: "ServicebusNamespaceAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusNamespaceOutput", GoGetter: "ServicebusNamespaceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusQueueAuthorizationRuleOutput", GoGetter: "ServicebusQueueAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusQueueOutput", GoGetter: "ServicebusQueueOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusSubscriptionOutput", GoGetter: "ServicebusSubscriptionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusSubscriptionRuleOutput", GoGetter: "ServicebusSubscriptionRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusTopicAuthorizationRuleOutput", GoGetter: "ServicebusTopicAuthorizationRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "servicebusTopicOutput", GoGetter: "ServicebusTopicOutput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceFabricClusterOutput", GoGetter: "ServiceFabricClusterOutput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedImageGalleryOutput", GoGetter: "SharedImageGalleryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedImageOutput", GoGetter: "SharedImageOutput"},
			_jsii_.MemberProperty{JsiiProperty: "signalrServiceOutput", GoGetter: "SignalrServiceOutput"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotsOutput", GoGetter: "SnapshotsOutput"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sqlElasticpoolOutput", GoGetter: "SqlElasticpoolOutput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlFailoverGroupOutput", GoGetter: "SqlFailoverGroupOutput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlFirewallRuleOutput", GoGetter: "SqlFirewallRuleOutput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerOutput", GoGetter: "SqlServerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountOutput", GoGetter: "StorageAccountOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageBlobOutput", GoGetter: "StorageBlobOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageContainerOutput", GoGetter: "StorageContainerOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageDataLakeGen2FilesystemOutput", GoGetter: "StorageDataLakeGen2FilesystemOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageQueueOutput", GoGetter: "StorageQueueOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageShareDirectoryOutput", GoGetter: "StorageShareDirectoryOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageShareOutput", GoGetter: "StorageShareOutput"},
			_jsii_.MemberProperty{JsiiProperty: "storageTableOutput", GoGetter: "StorageTableOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsFunctionJavascriptUdfOutput", GoGetter: "StreamAnalyticsFunctionJavascriptUdfOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsJobOutput", GoGetter: "StreamAnalyticsJobOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsOutputBlobOutput", GoGetter: "StreamAnalyticsOutputBlobOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsOutputEventhubOutput", GoGetter: "StreamAnalyticsOutputEventhubOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsOutputMssqlOutput", GoGetter: "StreamAnalyticsOutputMssqlOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsOutputServicebusQueueOutput", GoGetter: "StreamAnalyticsOutputServicebusQueueOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsOutputServicebusTopicOutput", GoGetter: "StreamAnalyticsOutputServicebusTopicOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsReferenceInputBlobOutput", GoGetter: "StreamAnalyticsReferenceInputBlobOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsStreamInputBlobOutput", GoGetter: "StreamAnalyticsStreamInputBlobOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsStreamInputEventhubOutput", GoGetter: "StreamAnalyticsStreamInputEventhubOutput"},
			_jsii_.MemberProperty{JsiiProperty: "streamAnalyticsStreamInputIothubOutput", GoGetter: "StreamAnalyticsStreamInputIothubOutput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetOutput", GoGetter: "SubnetOutput"},
			_jsii_.MemberProperty{JsiiProperty: "suffix", GoGetter: "Suffix"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "templateDeploymentOutput", GoGetter: "TemplateDeploymentOutput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trafficManagerProfileOutput", GoGetter: "TrafficManagerProfileOutput"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueIncludeNumbers", GoGetter: "UniqueIncludeNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueLength", GoGetter: "UniqueLength"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueSeed", GoGetter: "UniqueSeed"},
			_jsii_.MemberProperty{JsiiProperty: "uniqueSeedOutput", GoGetter: "UniqueSeedOutput"},
			_jsii_.MemberProperty{JsiiProperty: "userAssignedIdentityOutput", GoGetter: "UserAssignedIdentityOutput"},
			_jsii_.MemberProperty{JsiiProperty: "validationOutput", GoGetter: "ValidationOutput"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineExtensionOutput", GoGetter: "VirtualMachineExtensionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineOutput", GoGetter: "VirtualMachineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineScaleSetExtensionOutput", GoGetter: "VirtualMachineScaleSetExtensionOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualMachineScaleSetOutput", GoGetter: "VirtualMachineScaleSetOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkGatewayOutput", GoGetter: "VirtualNetworkGatewayOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkOutput", GoGetter: "VirtualNetworkOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkPeeringOutput", GoGetter: "VirtualNetworkPeeringOutput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualWanOutput", GoGetter: "VirtualWanOutput"},
			_jsii_.MemberProperty{JsiiProperty: "windowsVirtualMachineOutput", GoGetter: "WindowsVirtualMachineOutput"},
			_jsii_.MemberProperty{JsiiProperty: "windowsVirtualMachineScaleSetOutput", GoGetter: "WindowsVirtualMachineScaleSetOutput"},
		},
		func() interface{} {
			j := jsiiProxy_Naming{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"naming.NamingConfig",
		reflect.TypeOf((*NamingConfig)(nil)).Elem(),
	)
}