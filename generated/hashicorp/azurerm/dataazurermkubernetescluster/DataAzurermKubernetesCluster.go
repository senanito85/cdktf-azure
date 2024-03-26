package dataazurermkubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/dataazurermkubernetescluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/kubernetes_cluster azurerm_kubernetes_cluster}.
type DataAzurermKubernetesCluster interface {
	cdktf.TerraformDataSource
	AciConnectorLinux() DataAzurermKubernetesClusterAciConnectorLinuxList
	AddonProfile() DataAzurermKubernetesClusterAddonProfileList
	AgentPoolProfile() DataAzurermKubernetesClusterAgentPoolProfileList
	ApiServerAuthorizedIpRanges() *[]*string
	AzureActiveDirectoryRoleBasedAccessControl() DataAzurermKubernetesClusterAzureActiveDirectoryRoleBasedAccessControlList
	AzurePolicyEnabled() cdktf.IResolvable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionSetId() *string
	DnsPrefix() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpApplicationRoutingEnabled() cdktf.IResolvable
	HttpApplicationRoutingZoneName() *string
	Id() *string
	SetId(val *string)
	Identity() DataAzurermKubernetesClusterIdentityList
	IdInput() *string
	IngressApplicationGateway() DataAzurermKubernetesClusterIngressApplicationGatewayList
	KeyVaultSecretsProvider() DataAzurermKubernetesClusterKeyVaultSecretsProviderList
	KubeAdminConfig() DataAzurermKubernetesClusterKubeAdminConfigList
	KubeAdminConfigRaw() *string
	KubeConfig() DataAzurermKubernetesClusterKubeConfigList
	KubeConfigRaw() *string
	KubeletIdentity() DataAzurermKubernetesClusterKubeletIdentityList
	KubernetesVersion() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinuxProfile() DataAzurermKubernetesClusterLinuxProfileList
	Location() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkProfile() DataAzurermKubernetesClusterNetworkProfileList
	// The tree node.
	Node() constructs.Node
	NodeResourceGroup() *string
	OmsAgent() DataAzurermKubernetesClusterOmsAgentList
	OpenServiceMeshEnabled() cdktf.IResolvable
	PrivateClusterEnabled() cdktf.IResolvable
	PrivateFqdn() *string
	PrivateLinkEnabled() cdktf.IResolvable
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RoleBasedAccessControl() DataAzurermKubernetesClusterRoleBasedAccessControlList
	RoleBasedAccessControlEnabled() cdktf.IResolvable
	ServicePrincipal() DataAzurermKubernetesClusterServicePrincipalList
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermKubernetesClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WindowsProfile() DataAzurermKubernetesClusterWindowsProfileList
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DataAzurermKubernetesClusterTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAzurermKubernetesCluster
type jsiiProxy_DataAzurermKubernetesCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) AciConnectorLinux() DataAzurermKubernetesClusterAciConnectorLinuxList {
	var returns DataAzurermKubernetesClusterAciConnectorLinuxList
	_jsii_.Get(
		j,
		"aciConnectorLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) AddonProfile() DataAzurermKubernetesClusterAddonProfileList {
	var returns DataAzurermKubernetesClusterAddonProfileList
	_jsii_.Get(
		j,
		"addonProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) AgentPoolProfile() DataAzurermKubernetesClusterAgentPoolProfileList {
	var returns DataAzurermKubernetesClusterAgentPoolProfileList
	_jsii_.Get(
		j,
		"agentPoolProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ApiServerAuthorizedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiServerAuthorizedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) AzureActiveDirectoryRoleBasedAccessControl() DataAzurermKubernetesClusterAzureActiveDirectoryRoleBasedAccessControlList {
	var returns DataAzurermKubernetesClusterAzureActiveDirectoryRoleBasedAccessControlList
	_jsii_.Get(
		j,
		"azureActiveDirectoryRoleBasedAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) AzurePolicyEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"azurePolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) DnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) HttpApplicationRoutingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"httpApplicationRoutingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) HttpApplicationRoutingZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApplicationRoutingZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Identity() DataAzurermKubernetesClusterIdentityList {
	var returns DataAzurermKubernetesClusterIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) IngressApplicationGateway() DataAzurermKubernetesClusterIngressApplicationGatewayList {
	var returns DataAzurermKubernetesClusterIngressApplicationGatewayList
	_jsii_.Get(
		j,
		"ingressApplicationGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KeyVaultSecretsProvider() DataAzurermKubernetesClusterKeyVaultSecretsProviderList {
	var returns DataAzurermKubernetesClusterKeyVaultSecretsProviderList
	_jsii_.Get(
		j,
		"keyVaultSecretsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubeAdminConfig() DataAzurermKubernetesClusterKubeAdminConfigList {
	var returns DataAzurermKubernetesClusterKubeAdminConfigList
	_jsii_.Get(
		j,
		"kubeAdminConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubeAdminConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeAdminConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubeConfig() DataAzurermKubernetesClusterKubeConfigList {
	var returns DataAzurermKubernetesClusterKubeConfigList
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubeConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubeletIdentity() DataAzurermKubernetesClusterKubeletIdentityList {
	var returns DataAzurermKubernetesClusterKubeletIdentityList
	_jsii_.Get(
		j,
		"kubeletIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) LinuxProfile() DataAzurermKubernetesClusterLinuxProfileList {
	var returns DataAzurermKubernetesClusterLinuxProfileList
	_jsii_.Get(
		j,
		"linuxProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) NetworkProfile() DataAzurermKubernetesClusterNetworkProfileList {
	var returns DataAzurermKubernetesClusterNetworkProfileList
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) NodeResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) OmsAgent() DataAzurermKubernetesClusterOmsAgentList {
	var returns DataAzurermKubernetesClusterOmsAgentList
	_jsii_.Get(
		j,
		"omsAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) OpenServiceMeshEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"openServiceMeshEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) PrivateClusterEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"privateClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) PrivateFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) PrivateLinkEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"privateLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) RoleBasedAccessControl() DataAzurermKubernetesClusterRoleBasedAccessControlList {
	var returns DataAzurermKubernetesClusterRoleBasedAccessControlList
	_jsii_.Get(
		j,
		"roleBasedAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) RoleBasedAccessControlEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"roleBasedAccessControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) ServicePrincipal() DataAzurermKubernetesClusterServicePrincipalList {
	var returns DataAzurermKubernetesClusterServicePrincipalList
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) Timeouts() DataAzurermKubernetesClusterTimeoutsOutputReference {
	var returns DataAzurermKubernetesClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermKubernetesCluster) WindowsProfile() DataAzurermKubernetesClusterWindowsProfileList {
	var returns DataAzurermKubernetesClusterWindowsProfileList
	_jsii_.Get(
		j,
		"windowsProfile",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/kubernetes_cluster azurerm_kubernetes_cluster} Data Source.
func NewDataAzurermKubernetesCluster(scope constructs.Construct, id *string, config *DataAzurermKubernetesClusterConfig) DataAzurermKubernetesCluster {
	_init_.Initialize()

	if err := validateNewDataAzurermKubernetesClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermKubernetesCluster{}

	_jsii_.Create(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/kubernetes_cluster azurerm_kubernetes_cluster} Data Source.
func NewDataAzurermKubernetesCluster_Override(d DataAzurermKubernetesCluster, scope constructs.Construct, id *string, config *DataAzurermKubernetesClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermKubernetesCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermKubernetesCluster resource upon running "cdktf plan <stack-name>".
func DataAzurermKubernetesCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermKubernetesCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func DataAzurermKubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermKubernetesCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermKubernetesCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermKubernetesCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermKubernetesCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermKubernetesCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermKubernetesCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"azurerm.dataAzurermKubernetesCluster.DataAzurermKubernetesCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) PutTimeouts(value *DataAzurermKubernetesClusterTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermKubernetesCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

