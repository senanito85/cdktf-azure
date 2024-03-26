package naming

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/naming/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/naming/internal"
)

// Defines an Naming based on a Terraform module.
//
// Docs at Terraform Registry: {@link https://registry.terraform.io/modules/Azure/naming/azurerm/0.1.1 Azure/naming/azurerm}
type Naming interface {
	cdktf.TerraformModule
	AnalysisServicesServerOutput() *string
	ApiManagementOutput() *string
	AppConfigurationOutput() *string
	ApplicationGatewayOutput() *string
	ApplicationInsightsOutput() *string
	ApplicationSecurityGroupOutput() *string
	AppServiceOutput() *string
	AppServicePlanOutput() *string
	AutomationAccountOutput() *string
	AutomationCertificateOutput() *string
	AutomationCredentialOutput() *string
	AutomationRunbookOutput() *string
	AutomationScheduleOutput() *string
	AutomationVariableOutput() *string
	AvailabilitySetOutput() *string
	BastionHostOutput() *string
	BatchAccountOutput() *string
	BatchApplicationOutput() *string
	BatchCertificateOutput() *string
	BatchPoolOutput() *string
	BotChannelDirectlineOutput() *string
	BotChannelEmailOutput() *string
	BotChannelMsTeamsOutput() *string
	BotChannelSlackOutput() *string
	BotChannelsRegistrationOutput() *string
	BotConnectionOutput() *string
	BotWebAppOutput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnEndpointOutput() *string
	CdnProfileOutput() *string
	CognitiveAccountOutput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerGroupOutput() *string
	ContainerRegistryOutput() *string
	ContainerRegistryWebhookOutput() *string
	CosmosdbAccountOutput() *string
	CustomProviderOutput() *string
	DashboardOutput() *string
	DatabaseMigrationProjectOutput() *string
	DatabaseMigrationServiceOutput() *string
	DatabricksClusterOutput() *string
	DatabricksHighConcurrencyClusterOutput() *string
	DatabricksStandardClusterOutput() *string
	DatabricksWorkspaceOutput() *string
	DataFactoryDatasetMysqlOutput() *string
	DataFactoryDatasetPostgresqlOutput() *string
	DataFactoryDatasetSqlServerTableOutput() *string
	DataFactoryIntegrationRuntimeManagedOutput() *string
	DataFactoryLinkedServiceDataLakeStorageGen2Output() *string
	DataFactoryLinkedServiceKeyVaultOutput() *string
	DataFactoryLinkedServiceMysqlOutput() *string
	DataFactoryLinkedServicePostgresqlOutput() *string
	DataFactoryLinkedServiceSqlServerOutput() *string
	DataFactoryOutput() *string
	DataFactoryPipelineOutput() *string
	DataFactoryTriggerScheduleOutput() *string
	DataLakeAnalyticsAccountOutput() *string
	DataLakeAnalyticsFirewallRuleOutput() *string
	DataLakeStoreFirewallRuleOutput() *string
	DataLakeStoreOutput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DevTestLabOutput() *string
	DevTestLinuxVirtualMachineOutput() *string
	DevTestWindowsVirtualMachineOutput() *string
	DiskEncryptionSetOutput() *string
	DnsAaaaRecordOutput() *string
	DnsARecordOutput() *string
	DnsCaaRecordOutput() *string
	DnsCnameRecordOutput() *string
	DnsMxRecordOutput() *string
	DnsNsRecordOutput() *string
	DnsPtrRecordOutput() *string
	DnsTxtRecordOutput() *string
	DnsZoneOutput() *string
	EventgridDomainOutput() *string
	EventgridDomainTopicOutput() *string
	EventgridEventSubscriptionOutput() *string
	EventgridTopicOutput() *string
	EventhubAuthorizationRuleOutput() *string
	EventhubConsumerGroupOutput() *string
	EventhubNamespaceAuthorizationRuleOutput() *string
	EventhubNamespaceDisasterRecoveryConfigOutput() *string
	EventhubNamespaceOutput() *string
	EventhubOutput() *string
	ExpressRouteCircuitOutput() *string
	ExpressRouteGatewayOutput() *string
	FirewallApplicationRuleCollectionOutput() *string
	FirewallIpConfigurationOutput() *string
	FirewallNatRuleCollectionOutput() *string
	FirewallNetworkRuleCollectionOutput() *string
	FirewallOutput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontdoorFirewallPolicyOutput() *string
	FrontdoorOutput() *string
	FunctionAppOutput() *string
	HdinsightHadoopClusterOutput() *string
	HdinsightHbaseClusterOutput() *string
	HdinsightInteractiveQueryClusterOutput() *string
	HdinsightKafkaClusterOutput() *string
	HdinsightMlServicesClusterOutput() *string
	HdinsightRserverClusterOutput() *string
	HdinsightSparkClusterOutput() *string
	HdinsightStormClusterOutput() *string
	ImageOutput() *string
	IotcentralApplicationOutput() *string
	IothubConsumerGroupOutput() *string
	IothubDpsCertificateOutput() *string
	IothubDpsOutput() *string
	IothubOutput() *string
	KeyVaultCertificateOutput() *string
	KeyVaultKeyOutput() *string
	KeyVaultOutput() *string
	KeyVaultSecretOutput() *string
	KubernetesClusterOutput() *string
	KustoClusterOutput() *string
	KustoDatabaseOutput() *string
	KustoEventhubDataConnectionOutput() *string
	LbNatRuleOutput() *string
	LbOutput() *string
	LinuxVirtualMachineOutput() *string
	LinuxVirtualMachineScaleSetOutput() *string
	LocalNetworkGatewayOutput() *string
	LogAnalyticsWorkspaceOutput() *string
	MachineLearningWorkspaceOutput() *string
	ManagedDiskOutput() *string
	MapsAccountOutput() *string
	MariadbDatabaseOutput() *string
	MariadbFirewallRuleOutput() *string
	MariadbServerOutput() *string
	MariadbVirtualNetworkRuleOutput() *string
	MssqlDatabaseOutput() *string
	MssqlElasticpoolOutput() *string
	MssqlServerOutput() *string
	MysqlDatabaseOutput() *string
	MysqlFirewallRuleOutput() *string
	MysqlServerOutput() *string
	MysqlVirtualNetworkRuleOutput() *string
	NetworkDdosProtectionPlanOutput() *string
	NetworkInterfaceOutput() *string
	NetworkSecurityGroupOutput() *string
	NetworkSecurityGroupRuleOutput() *string
	NetworkSecurityRuleOutput() *string
	NetworkWatcherOutput() *string
	// The tree node.
	Node() constructs.Node
	NotificationHubAuthorizationRuleOutput() *string
	NotificationHubNamespaceOutput() *string
	NotificationHubOutput() *string
	PointToSiteVpnGatewayOutput() *string
	PostgresqlDatabaseOutput() *string
	PostgresqlFirewallRuleOutput() *string
	PostgresqlServerOutput() *string
	PostgresqlVirtualNetworkRuleOutput() *string
	PowerbiEmbeddedOutput() *string
	Prefix() *[]*string
	SetPrefix(val *[]*string)
	PrivateDnsAaaaRecordOutput() *string
	PrivateDnsARecordOutput() *string
	PrivateDnsCnameRecordOutput() *string
	PrivateDnsMxRecordOutput() *string
	PrivateDnsPtrRecordOutput() *string
	PrivateDnsSrvRecordOutput() *string
	PrivateDnsTxtRecordOutput() *string
	PrivateDnsZoneGroupOutput() *string
	PrivateDnsZoneOutput() *string
	PrivateEndpointOutput() *string
	PrivateLinkServiceOutput() *string
	PrivateServiceConnectionOutput() *string
	// Experimental.
	Providers() *[]interface{}
	ProximityPlacementGroupOutput() *string
	PublicIpOutput() *string
	PublicIpPrefixOutput() *string
	// Experimental.
	RawOverrides() interface{}
	RedisCacheOutput() *string
	RedisFirewallRuleOutput() *string
	RelayHybridConnectionOutput() *string
	RelayNamespaceOutput() *string
	ResourceGroupOutput() *string
	RoleAssignmentOutput() *string
	RoleDefinitionOutput() *string
	RouteOutput() *string
	RouteTableOutput() *string
	ServicebusNamespaceAuthorizationRuleOutput() *string
	ServicebusNamespaceOutput() *string
	ServicebusQueueAuthorizationRuleOutput() *string
	ServicebusQueueOutput() *string
	ServicebusSubscriptionOutput() *string
	ServicebusSubscriptionRuleOutput() *string
	ServicebusTopicAuthorizationRuleOutput() *string
	ServicebusTopicOutput() *string
	ServiceFabricClusterOutput() *string
	SharedImageGalleryOutput() *string
	SharedImageOutput() *string
	SignalrServiceOutput() *string
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	SnapshotsOutput() *string
	// Experimental.
	Source() *string
	SqlElasticpoolOutput() *string
	SqlFailoverGroupOutput() *string
	SqlFirewallRuleOutput() *string
	SqlServerOutput() *string
	StorageAccountOutput() *string
	StorageBlobOutput() *string
	StorageContainerOutput() *string
	StorageDataLakeGen2FilesystemOutput() *string
	StorageQueueOutput() *string
	StorageShareDirectoryOutput() *string
	StorageShareOutput() *string
	StorageTableOutput() *string
	StreamAnalyticsFunctionJavascriptUdfOutput() *string
	StreamAnalyticsJobOutput() *string
	StreamAnalyticsOutputBlobOutput() *string
	StreamAnalyticsOutputEventhubOutput() *string
	StreamAnalyticsOutputMssqlOutput() *string
	StreamAnalyticsOutputServicebusQueueOutput() *string
	StreamAnalyticsOutputServicebusTopicOutput() *string
	StreamAnalyticsReferenceInputBlobOutput() *string
	StreamAnalyticsStreamInputBlobOutput() *string
	StreamAnalyticsStreamInputEventhubOutput() *string
	StreamAnalyticsStreamInputIothubOutput() *string
	SubnetOutput() *string
	Suffix() *[]*string
	SetSuffix(val *[]*string)
	TemplateDeploymentOutput() *string
	TrafficManagerProfileOutput() *string
	UniqueIncludeNumbers() *bool
	SetUniqueIncludeNumbers(val *bool)
	UniqueLength() *float64
	SetUniqueLength(val *float64)
	UniqueSeed() *string
	SetUniqueSeed(val *string)
	UniqueSeedOutput() *string
	UserAssignedIdentityOutput() *string
	ValidationOutput() *string
	// Experimental.
	Version() *string
	VirtualMachineExtensionOutput() *string
	VirtualMachineOutput() *string
	VirtualMachineScaleSetExtensionOutput() *string
	VirtualMachineScaleSetOutput() *string
	VirtualNetworkGatewayOutput() *string
	VirtualNetworkOutput() *string
	VirtualNetworkPeeringOutput() *string
	VirtualWanOutput() *string
	WindowsVirtualMachineOutput() *string
	WindowsVirtualMachineScaleSetOutput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Naming
type jsiiProxy_Naming struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Naming) AnalysisServicesServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisServicesServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ApiManagementOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AppConfigurationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appConfigurationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ApplicationGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ApplicationInsightsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ApplicationSecurityGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSecurityGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AppServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AppServicePlanOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appServicePlanOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationCertificateOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationCertificateOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationCredentialOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationCredentialOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationRunbookOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationRunbookOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationScheduleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationScheduleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AutomationVariableOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationVariableOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) AvailabilitySetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BastionHostOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bastionHostOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BatchAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BatchApplicationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchApplicationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BatchCertificateOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchCertificateOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BatchPoolOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchPoolOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotChannelDirectlineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botChannelDirectlineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotChannelEmailOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botChannelEmailOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotChannelMsTeamsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botChannelMsTeamsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotChannelSlackOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botChannelSlackOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotChannelsRegistrationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botChannelsRegistrationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotConnectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botConnectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) BotWebAppOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botWebAppOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CdnEndpointOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnEndpointOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CdnProfileOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdnProfileOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CognitiveAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cognitiveAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ContainerGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ContainerRegistryOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ContainerRegistryWebhookOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRegistryWebhookOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CosmosdbAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) CustomProviderOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customProviderOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DashboardOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabaseMigrationProjectOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationProjectOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabaseMigrationServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseMigrationServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabricksClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabricksHighConcurrencyClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksHighConcurrencyClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabricksStandardClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksStandardClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DatabricksWorkspaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksWorkspaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryDatasetMysqlOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryDatasetMysqlOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryDatasetPostgresqlOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryDatasetPostgresqlOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryDatasetSqlServerTableOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryDatasetSqlServerTableOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryIntegrationRuntimeManagedOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIntegrationRuntimeManagedOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryLinkedServiceDataLakeStorageGen2Output() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryLinkedServiceDataLakeStorageGen2Output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryLinkedServiceKeyVaultOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryLinkedServiceKeyVaultOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryLinkedServiceMysqlOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryLinkedServiceMysqlOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryLinkedServicePostgresqlOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryLinkedServicePostgresqlOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryLinkedServiceSqlServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryLinkedServiceSqlServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryPipelineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryPipelineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataFactoryTriggerScheduleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryTriggerScheduleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataLakeAnalyticsAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeAnalyticsAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataLakeAnalyticsFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeAnalyticsFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataLakeStoreFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeStoreFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DataLakeStoreOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLakeStoreOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DevTestLabOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devTestLabOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DevTestLinuxVirtualMachineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devTestLinuxVirtualMachineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DevTestWindowsVirtualMachineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devTestWindowsVirtualMachineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DiskEncryptionSetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsAaaaRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsAaaaRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsARecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsARecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsCaaRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCaaRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsCnameRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCnameRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsMxRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsMxRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsNsRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNsRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsPtrRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPtrRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsTxtRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsTxtRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) DnsZoneOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsZoneOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventgridDomainOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridDomainOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventgridDomainTopicOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridDomainTopicOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventgridEventSubscriptionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridEventSubscriptionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventgridTopicOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventgridTopicOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubConsumerGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubNamespaceAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubNamespaceDisasterRecoveryConfigOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceDisasterRecoveryConfigOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubNamespaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNamespaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) EventhubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ExpressRouteCircuitOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ExpressRouteGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FirewallApplicationRuleCollectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallApplicationRuleCollectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FirewallIpConfigurationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallIpConfigurationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FirewallNatRuleCollectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallNatRuleCollectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FirewallNetworkRuleCollectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallNetworkRuleCollectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FirewallOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firewallOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FrontdoorFirewallPolicyOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontdoorFirewallPolicyOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FrontdoorOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontdoorOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) FunctionAppOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAppOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightHadoopClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightHadoopClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightHbaseClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightHbaseClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightInteractiveQueryClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightInteractiveQueryClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightKafkaClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightKafkaClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightMlServicesClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightMlServicesClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightRserverClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightRserverClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightSparkClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightSparkClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) HdinsightStormClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hdinsightStormClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ImageOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) IotcentralApplicationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iotcentralApplicationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) IothubConsumerGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubConsumerGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) IothubDpsCertificateOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubDpsCertificateOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) IothubDpsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubDpsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) IothubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KeyVaultCertificateOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultCertificateOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KeyVaultKeyOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KeyVaultOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KeyVaultSecretOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultSecretOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KubernetesClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KustoClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KustoDatabaseOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) KustoEventhubDataConnectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoEventhubDataConnectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LbNatRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbNatRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LbOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lbOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LinuxVirtualMachineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxVirtualMachineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LinuxVirtualMachineScaleSetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxVirtualMachineScaleSetOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LocalNetworkGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localNetworkGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) LogAnalyticsWorkspaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MachineLearningWorkspaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ManagedDiskOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDiskOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MapsAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mapsAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MariadbDatabaseOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbDatabaseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MariadbFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MariadbServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MariadbVirtualNetworkRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariadbVirtualNetworkRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MssqlDatabaseOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mssqlDatabaseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MssqlElasticpoolOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mssqlElasticpoolOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MssqlServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mssqlServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MysqlDatabaseOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlDatabaseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MysqlFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MysqlServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) MysqlVirtualNetworkRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlVirtualNetworkRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkDdosProtectionPlanOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkDdosProtectionPlanOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkInterfaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkSecurityGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkSecurityGroupRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkSecurityRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NetworkWatcherOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkWatcherOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NotificationHubAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationHubAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NotificationHubNamespaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationHubNamespaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) NotificationHubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationHubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PointToSiteVpnGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pointToSiteVpnGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PostgresqlDatabaseOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlDatabaseOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PostgresqlFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PostgresqlServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PostgresqlVirtualNetworkRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlVirtualNetworkRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PowerbiEmbeddedOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"powerbiEmbeddedOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Prefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsAaaaRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsAaaaRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsARecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsARecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsCnameRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsCnameRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsMxRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsMxRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsPtrRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPtrRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsSrvRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsSrvRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsTxtRecordOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsTxtRecordOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsZoneGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateDnsZoneOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateEndpointOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateLinkServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PrivateServiceConnectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateServiceConnectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ProximityPlacementGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PublicIpOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) PublicIpPrefixOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RedisCacheOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCacheOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RedisFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RelayHybridConnectionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayHybridConnectionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RelayNamespaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayNamespaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ResourceGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RoleAssignmentOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleAssignmentOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RoleDefinitionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleDefinitionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RouteOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) RouteTableOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusNamespaceAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespaceAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusNamespaceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespaceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusQueueAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusQueueAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusQueueOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusQueueOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusSubscriptionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusSubscriptionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusSubscriptionRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusSubscriptionRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusTopicAuthorizationRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusTopicAuthorizationRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServicebusTopicOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusTopicOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ServiceFabricClusterOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceFabricClusterOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SharedImageGalleryOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedImageGalleryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SharedImageOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedImageOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SignalrServiceOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signalrServiceOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SnapshotsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SqlElasticpoolOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlElasticpoolOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SqlFailoverGroupOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlFailoverGroupOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SqlFirewallRuleOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlFirewallRuleOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SqlServerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageAccountOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageBlobOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageContainerOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageDataLakeGen2FilesystemOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageDataLakeGen2FilesystemOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageQueueOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageQueueOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageShareDirectoryOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageShareDirectoryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageShareOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageShareOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StorageTableOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTableOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsFunctionJavascriptUdfOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsFunctionJavascriptUdfOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsJobOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsJobOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsOutputBlobOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsOutputBlobOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsOutputEventhubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsOutputEventhubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsOutputMssqlOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsOutputMssqlOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsOutputServicebusQueueOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsOutputServicebusQueueOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsOutputServicebusTopicOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsOutputServicebusTopicOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsReferenceInputBlobOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsReferenceInputBlobOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsStreamInputBlobOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsStreamInputBlobOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsStreamInputEventhubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsStreamInputEventhubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) StreamAnalyticsStreamInputIothubOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsStreamInputIothubOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) SubnetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Suffix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) TemplateDeploymentOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDeploymentOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) TrafficManagerProfileOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficManagerProfileOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) UniqueIncludeNumbers() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"uniqueIncludeNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) UniqueLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uniqueLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) UniqueSeed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueSeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) UniqueSeedOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueSeedOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) UserAssignedIdentityOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) ValidationOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualMachineExtensionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineExtensionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualMachineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualMachineScaleSetExtensionOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetExtensionOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualMachineScaleSetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualNetworkGatewayOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkGatewayOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualNetworkOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualNetworkPeeringOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkPeeringOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) VirtualWanOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualWanOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) WindowsVirtualMachineOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsVirtualMachineOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Naming) WindowsVirtualMachineScaleSetOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsVirtualMachineScaleSetOutput",
		&returns,
	)
	return returns
}


func NewNaming(scope constructs.Construct, id *string, config *NamingConfig) Naming {
	_init_.Initialize()

	if err := validateNewNamingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Naming{}

	_jsii_.Create(
		"naming.Naming",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewNaming_Override(n Naming, scope constructs.Construct, id *string, config *NamingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"naming.Naming",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_Naming)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Naming)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Naming)SetPrefix(val *[]*string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_Naming)SetSuffix(val *[]*string) {
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_Naming)SetUniqueIncludeNumbers(val *bool) {
	_jsii_.Set(
		j,
		"uniqueIncludeNumbers",
		val,
	)
}

func (j *jsiiProxy_Naming)SetUniqueLength(val *float64) {
	_jsii_.Set(
		j,
		"uniqueLength",
		val,
	)
}

func (j *jsiiProxy_Naming)SetUniqueSeed(val *string) {
	_jsii_.Set(
		j,
		"uniqueSeed",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Naming_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNaming_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"naming.Naming",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Naming_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNaming_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"naming.Naming",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_Naming) AddProvider(provider interface{}) {
	if err := n.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addProvider",
		[]interface{}{provider},
	)
}

func (n *jsiiProxy_Naming) GetString(output *string) *string {
	if err := n.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := n.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_Naming) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_Naming) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_Naming) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

