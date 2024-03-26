package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/kubernetescluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster azurerm_kubernetes_cluster}.
type KubernetesCluster interface {
	cdktf.TerraformResource
	AciConnectorLinux() KubernetesClusterAciConnectorLinuxOutputReference
	AciConnectorLinuxInput() *KubernetesClusterAciConnectorLinux
	AddonProfile() KubernetesClusterAddonProfileOutputReference
	AddonProfileInput() *KubernetesClusterAddonProfile
	ApiServerAuthorizedIpRanges() *[]*string
	SetApiServerAuthorizedIpRanges(val *[]*string)
	ApiServerAuthorizedIpRangesInput() *[]*string
	AutomaticChannelUpgrade() *string
	SetAutomaticChannelUpgrade(val *string)
	AutomaticChannelUpgradeInput() *string
	AutoScalerProfile() KubernetesClusterAutoScalerProfileOutputReference
	AutoScalerProfileInput() *KubernetesClusterAutoScalerProfile
	AzureActiveDirectoryRoleBasedAccessControl() KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference
	AzureActiveDirectoryRoleBasedAccessControlInput() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	AzurePolicyEnabled() interface{}
	SetAzurePolicyEnabled(val interface{})
	AzurePolicyEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultNodePool() KubernetesClusterDefaultNodePoolOutputReference
	DefaultNodePoolInput() *KubernetesClusterDefaultNodePool
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DnsPrefix() *string
	SetDnsPrefix(val *string)
	DnsPrefixInput() *string
	DnsPrefixPrivateCluster() *string
	SetDnsPrefixPrivateCluster(val *string)
	DnsPrefixPrivateClusterInput() *string
	EnablePodSecurityPolicy() interface{}
	SetEnablePodSecurityPolicy(val interface{})
	EnablePodSecurityPolicyInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Fqdn() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpApplicationRoutingEnabled() interface{}
	SetHttpApplicationRoutingEnabled(val interface{})
	HttpApplicationRoutingEnabledInput() interface{}
	HttpApplicationRoutingZoneName() *string
	HttpProxyConfig() KubernetesClusterHttpProxyConfigOutputReference
	HttpProxyConfigInput() *KubernetesClusterHttpProxyConfig
	Id() *string
	SetId(val *string)
	Identity() KubernetesClusterIdentityOutputReference
	IdentityInput() *KubernetesClusterIdentity
	IdInput() *string
	IngressApplicationGateway() KubernetesClusterIngressApplicationGatewayOutputReference
	IngressApplicationGatewayInput() *KubernetesClusterIngressApplicationGateway
	KeyVaultSecretsProvider() KubernetesClusterKeyVaultSecretsProviderOutputReference
	KeyVaultSecretsProviderInput() *KubernetesClusterKeyVaultSecretsProvider
	KubeAdminConfig() KubernetesClusterKubeAdminConfigList
	KubeAdminConfigRaw() *string
	KubeConfig() KubernetesClusterKubeConfigList
	KubeConfigRaw() *string
	KubeletIdentity() KubernetesClusterKubeletIdentityOutputReference
	KubeletIdentityInput() *KubernetesClusterKubeletIdentity
	KubernetesVersion() *string
	SetKubernetesVersion(val *string)
	KubernetesVersionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinuxProfile() KubernetesClusterLinuxProfileOutputReference
	LinuxProfileInput() *KubernetesClusterLinuxProfile
	LocalAccountDisabled() interface{}
	SetLocalAccountDisabled(val interface{})
	LocalAccountDisabledInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceWindow() KubernetesClusterMaintenanceWindowOutputReference
	MaintenanceWindowInput() *KubernetesClusterMaintenanceWindow
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkProfile() KubernetesClusterNetworkProfileOutputReference
	NetworkProfileInput() *KubernetesClusterNetworkProfile
	// The tree node.
	Node() constructs.Node
	NodeResourceGroup() *string
	SetNodeResourceGroup(val *string)
	NodeResourceGroupInput() *string
	OmsAgent() KubernetesClusterOmsAgentOutputReference
	OmsAgentInput() *KubernetesClusterOmsAgent
	OpenServiceMeshEnabled() interface{}
	SetOpenServiceMeshEnabled(val interface{})
	OpenServiceMeshEnabledInput() interface{}
	PortalFqdn() *string
	PrivateClusterEnabled() interface{}
	SetPrivateClusterEnabled(val interface{})
	PrivateClusterEnabledInput() interface{}
	PrivateClusterPublicFqdnEnabled() interface{}
	SetPrivateClusterPublicFqdnEnabled(val interface{})
	PrivateClusterPublicFqdnEnabledInput() interface{}
	PrivateDnsZoneId() *string
	SetPrivateDnsZoneId(val *string)
	PrivateDnsZoneIdInput() *string
	PrivateFqdn() *string
	PrivateLinkEnabled() interface{}
	SetPrivateLinkEnabled(val interface{})
	PrivateLinkEnabledInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RoleBasedAccessControl() KubernetesClusterRoleBasedAccessControlOutputReference
	RoleBasedAccessControlEnabled() interface{}
	SetRoleBasedAccessControlEnabled(val interface{})
	RoleBasedAccessControlEnabledInput() interface{}
	RoleBasedAccessControlInput() *KubernetesClusterRoleBasedAccessControl
	ServicePrincipal() KubernetesClusterServicePrincipalOutputReference
	ServicePrincipalInput() *KubernetesClusterServicePrincipal
	SkuTier() *string
	SetSkuTier(val *string)
	SkuTierInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WindowsProfile() KubernetesClusterWindowsProfileOutputReference
	WindowsProfileInput() *KubernetesClusterWindowsProfile
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAciConnectorLinux(value *KubernetesClusterAciConnectorLinux)
	PutAddonProfile(value *KubernetesClusterAddonProfile)
	PutAutoScalerProfile(value *KubernetesClusterAutoScalerProfile)
	PutAzureActiveDirectoryRoleBasedAccessControl(value *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl)
	PutDefaultNodePool(value *KubernetesClusterDefaultNodePool)
	PutHttpProxyConfig(value *KubernetesClusterHttpProxyConfig)
	PutIdentity(value *KubernetesClusterIdentity)
	PutIngressApplicationGateway(value *KubernetesClusterIngressApplicationGateway)
	PutKeyVaultSecretsProvider(value *KubernetesClusterKeyVaultSecretsProvider)
	PutKubeletIdentity(value *KubernetesClusterKubeletIdentity)
	PutLinuxProfile(value *KubernetesClusterLinuxProfile)
	PutMaintenanceWindow(value *KubernetesClusterMaintenanceWindow)
	PutNetworkProfile(value *KubernetesClusterNetworkProfile)
	PutOmsAgent(value *KubernetesClusterOmsAgent)
	PutRoleBasedAccessControl(value *KubernetesClusterRoleBasedAccessControl)
	PutServicePrincipal(value *KubernetesClusterServicePrincipal)
	PutTimeouts(value *KubernetesClusterTimeouts)
	PutWindowsProfile(value *KubernetesClusterWindowsProfile)
	ResetAciConnectorLinux()
	ResetAddonProfile()
	ResetApiServerAuthorizedIpRanges()
	ResetAutomaticChannelUpgrade()
	ResetAutoScalerProfile()
	ResetAzureActiveDirectoryRoleBasedAccessControl()
	ResetAzurePolicyEnabled()
	ResetDiskEncryptionSetId()
	ResetDnsPrefix()
	ResetDnsPrefixPrivateCluster()
	ResetEnablePodSecurityPolicy()
	ResetHttpApplicationRoutingEnabled()
	ResetHttpProxyConfig()
	ResetId()
	ResetIdentity()
	ResetIngressApplicationGateway()
	ResetKeyVaultSecretsProvider()
	ResetKubeletIdentity()
	ResetKubernetesVersion()
	ResetLinuxProfile()
	ResetLocalAccountDisabled()
	ResetMaintenanceWindow()
	ResetNetworkProfile()
	ResetNodeResourceGroup()
	ResetOmsAgent()
	ResetOpenServiceMeshEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateClusterEnabled()
	ResetPrivateClusterPublicFqdnEnabled()
	ResetPrivateDnsZoneId()
	ResetPrivateLinkEnabled()
	ResetPublicNetworkAccessEnabled()
	ResetRoleBasedAccessControl()
	ResetRoleBasedAccessControlEnabled()
	ResetServicePrincipal()
	ResetSkuTier()
	ResetTags()
	ResetTimeouts()
	ResetWindowsProfile()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for KubernetesCluster
type jsiiProxy_KubernetesCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesCluster) AciConnectorLinux() KubernetesClusterAciConnectorLinuxOutputReference {
	var returns KubernetesClusterAciConnectorLinuxOutputReference
	_jsii_.Get(
		j,
		"aciConnectorLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AciConnectorLinuxInput() *KubernetesClusterAciConnectorLinux {
	var returns *KubernetesClusterAciConnectorLinux
	_jsii_.Get(
		j,
		"aciConnectorLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AddonProfile() KubernetesClusterAddonProfileOutputReference {
	var returns KubernetesClusterAddonProfileOutputReference
	_jsii_.Get(
		j,
		"addonProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AddonProfileInput() *KubernetesClusterAddonProfile {
	var returns *KubernetesClusterAddonProfile
	_jsii_.Get(
		j,
		"addonProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ApiServerAuthorizedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiServerAuthorizedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ApiServerAuthorizedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiServerAuthorizedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutomaticChannelUpgrade() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticChannelUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutomaticChannelUpgradeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticChannelUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutoScalerProfile() KubernetesClusterAutoScalerProfileOutputReference {
	var returns KubernetesClusterAutoScalerProfileOutputReference
	_jsii_.Get(
		j,
		"autoScalerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AutoScalerProfileInput() *KubernetesClusterAutoScalerProfile {
	var returns *KubernetesClusterAutoScalerProfile
	_jsii_.Get(
		j,
		"autoScalerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzureActiveDirectoryRoleBasedAccessControl() KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference {
	var returns KubernetesClusterAzureActiveDirectoryRoleBasedAccessControlOutputReference
	_jsii_.Get(
		j,
		"azureActiveDirectoryRoleBasedAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzureActiveDirectoryRoleBasedAccessControlInput() *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl {
	var returns *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl
	_jsii_.Get(
		j,
		"azureActiveDirectoryRoleBasedAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzurePolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azurePolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) AzurePolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azurePolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DefaultNodePool() KubernetesClusterDefaultNodePoolOutputReference {
	var returns KubernetesClusterDefaultNodePoolOutputReference
	_jsii_.Get(
		j,
		"defaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DefaultNodePoolInput() *KubernetesClusterDefaultNodePool {
	var returns *KubernetesClusterDefaultNodePool
	_jsii_.Get(
		j,
		"defaultNodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixPrivateCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixPrivateCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) DnsPrefixPrivateClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrefixPrivateClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EnablePodSecurityPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePodSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) EnablePodSecurityPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePodSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpApplicationRoutingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpApplicationRoutingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpApplicationRoutingZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApplicationRoutingZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpProxyConfig() KubernetesClusterHttpProxyConfigOutputReference {
	var returns KubernetesClusterHttpProxyConfigOutputReference
	_jsii_.Get(
		j,
		"httpProxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) HttpProxyConfigInput() *KubernetesClusterHttpProxyConfig {
	var returns *KubernetesClusterHttpProxyConfig
	_jsii_.Get(
		j,
		"httpProxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Identity() KubernetesClusterIdentityOutputReference {
	var returns KubernetesClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdentityInput() *KubernetesClusterIdentity {
	var returns *KubernetesClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IngressApplicationGateway() KubernetesClusterIngressApplicationGatewayOutputReference {
	var returns KubernetesClusterIngressApplicationGatewayOutputReference
	_jsii_.Get(
		j,
		"ingressApplicationGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) IngressApplicationGatewayInput() *KubernetesClusterIngressApplicationGateway {
	var returns *KubernetesClusterIngressApplicationGateway
	_jsii_.Get(
		j,
		"ingressApplicationGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KeyVaultSecretsProvider() KubernetesClusterKeyVaultSecretsProviderOutputReference {
	var returns KubernetesClusterKeyVaultSecretsProviderOutputReference
	_jsii_.Get(
		j,
		"keyVaultSecretsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KeyVaultSecretsProviderInput() *KubernetesClusterKeyVaultSecretsProvider {
	var returns *KubernetesClusterKeyVaultSecretsProvider
	_jsii_.Get(
		j,
		"keyVaultSecretsProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeAdminConfig() KubernetesClusterKubeAdminConfigList {
	var returns KubernetesClusterKubeAdminConfigList
	_jsii_.Get(
		j,
		"kubeAdminConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeAdminConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeAdminConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeConfig() KubernetesClusterKubeConfigList {
	var returns KubernetesClusterKubeConfigList
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeletIdentity() KubernetesClusterKubeletIdentityOutputReference {
	var returns KubernetesClusterKubeletIdentityOutputReference
	_jsii_.Get(
		j,
		"kubeletIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubeletIdentityInput() *KubernetesClusterKubeletIdentity {
	var returns *KubernetesClusterKubeletIdentity
	_jsii_.Get(
		j,
		"kubeletIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) KubernetesVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LinuxProfile() KubernetesClusterLinuxProfileOutputReference {
	var returns KubernetesClusterLinuxProfileOutputReference
	_jsii_.Get(
		j,
		"linuxProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LinuxProfileInput() *KubernetesClusterLinuxProfile {
	var returns *KubernetesClusterLinuxProfile
	_jsii_.Get(
		j,
		"linuxProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocalAccountDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAccountDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocalAccountDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAccountDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenanceWindow() KubernetesClusterMaintenanceWindowOutputReference {
	var returns KubernetesClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) MaintenanceWindowInput() *KubernetesClusterMaintenanceWindow {
	var returns *KubernetesClusterMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkProfile() KubernetesClusterNetworkProfileOutputReference {
	var returns KubernetesClusterNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NetworkProfileInput() *KubernetesClusterNetworkProfile {
	var returns *KubernetesClusterNetworkProfile
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) NodeResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OmsAgent() KubernetesClusterOmsAgentOutputReference {
	var returns KubernetesClusterOmsAgentOutputReference
	_jsii_.Get(
		j,
		"omsAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OmsAgentInput() *KubernetesClusterOmsAgent {
	var returns *KubernetesClusterOmsAgent
	_jsii_.Get(
		j,
		"omsAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OpenServiceMeshEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openServiceMeshEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) OpenServiceMeshEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openServiceMeshEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PortalFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterPublicFqdnEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterPublicFqdnEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateClusterPublicFqdnEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateClusterPublicFqdnEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateDnsZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateDnsZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateLinkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PrivateLinkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControl() KubernetesClusterRoleBasedAccessControlOutputReference {
	var returns KubernetesClusterRoleBasedAccessControlOutputReference
	_jsii_.Get(
		j,
		"roleBasedAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleBasedAccessControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleBasedAccessControlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) RoleBasedAccessControlInput() *KubernetesClusterRoleBasedAccessControl {
	var returns *KubernetesClusterRoleBasedAccessControl
	_jsii_.Get(
		j,
		"roleBasedAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServicePrincipal() KubernetesClusterServicePrincipalOutputReference {
	var returns KubernetesClusterServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) ServicePrincipalInput() *KubernetesClusterServicePrincipal {
	var returns *KubernetesClusterServicePrincipal
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SkuTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) SkuTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) Timeouts() KubernetesClusterTimeoutsOutputReference {
	var returns KubernetesClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) WindowsProfile() KubernetesClusterWindowsProfileOutputReference {
	var returns KubernetesClusterWindowsProfileOutputReference
	_jsii_.Get(
		j,
		"windowsProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesCluster) WindowsProfileInput() *KubernetesClusterWindowsProfile {
	var returns *KubernetesClusterWindowsProfile
	_jsii_.Get(
		j,
		"windowsProfileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster azurerm_kubernetes_cluster} Resource.
func NewKubernetesCluster(scope constructs.Construct, id *string, config *KubernetesClusterConfig) KubernetesCluster {
	_init_.Initialize()

	if err := validateNewKubernetesClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesCluster{}

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster azurerm_kubernetes_cluster} Resource.
func NewKubernetesCluster_Override(k KubernetesCluster, scope constructs.Construct, id *string, config *KubernetesClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetApiServerAuthorizedIpRanges(val *[]*string) {
	if err := j.validateSetApiServerAuthorizedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiServerAuthorizedIpRanges",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetAutomaticChannelUpgrade(val *string) {
	if err := j.validateSetAutomaticChannelUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticChannelUpgrade",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetAzurePolicyEnabled(val interface{}) {
	if err := j.validateSetAzurePolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azurePolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDnsPrefix(val *string) {
	if err := j.validateSetDnsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPrefix",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetDnsPrefixPrivateCluster(val *string) {
	if err := j.validateSetDnsPrefixPrivateClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPrefixPrivateCluster",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetEnablePodSecurityPolicy(val interface{}) {
	if err := j.validateSetEnablePodSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePodSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetHttpApplicationRoutingEnabled(val interface{}) {
	if err := j.validateSetHttpApplicationRoutingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpApplicationRoutingEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetKubernetesVersion(val *string) {
	if err := j.validateSetKubernetesVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetLocalAccountDisabled(val interface{}) {
	if err := j.validateSetLocalAccountDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAccountDisabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetNodeResourceGroup(val *string) {
	if err := j.validateSetNodeResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeResourceGroup",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetOpenServiceMeshEnabled(val interface{}) {
	if err := j.validateSetOpenServiceMeshEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openServiceMeshEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetPrivateClusterEnabled(val interface{}) {
	if err := j.validateSetPrivateClusterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetPrivateClusterPublicFqdnEnabled(val interface{}) {
	if err := j.validateSetPrivateClusterPublicFqdnEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateClusterPublicFqdnEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetPrivateDnsZoneId(val *string) {
	if err := j.validateSetPrivateDnsZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsZoneId",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetPrivateLinkEnabled(val interface{}) {
	if err := j.validateSetPrivateLinkEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetRoleBasedAccessControlEnabled(val interface{}) {
	if err := j.validateSetRoleBasedAccessControlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleBasedAccessControlEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetSkuTier(val *string) {
	if err := j.validateSetSkuTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuTier",
		val,
	)
}

func (j *jsiiProxy_KubernetesCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a KubernetesCluster resource upon running "cdktf plan <stack-name>".
func KubernetesCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"azurerm.kubernetesCluster.KubernetesCluster",
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
func KubernetesCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.kubernetesCluster.KubernetesCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.kubernetesCluster.KubernetesCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"azurerm.kubernetesCluster.KubernetesCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"azurerm.kubernetesCluster.KubernetesCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesCluster) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesCluster) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesCluster) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesCluster) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAciConnectorLinux(value *KubernetesClusterAciConnectorLinux) {
	if err := k.validatePutAciConnectorLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAciConnectorLinux",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAddonProfile(value *KubernetesClusterAddonProfile) {
	if err := k.validatePutAddonProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAddonProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAutoScalerProfile(value *KubernetesClusterAutoScalerProfile) {
	if err := k.validatePutAutoScalerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAutoScalerProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutAzureActiveDirectoryRoleBasedAccessControl(value *KubernetesClusterAzureActiveDirectoryRoleBasedAccessControl) {
	if err := k.validatePutAzureActiveDirectoryRoleBasedAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAzureActiveDirectoryRoleBasedAccessControl",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutDefaultNodePool(value *KubernetesClusterDefaultNodePool) {
	if err := k.validatePutDefaultNodePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDefaultNodePool",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutHttpProxyConfig(value *KubernetesClusterHttpProxyConfig) {
	if err := k.validatePutHttpProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putHttpProxyConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutIdentity(value *KubernetesClusterIdentity) {
	if err := k.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutIngressApplicationGateway(value *KubernetesClusterIngressApplicationGateway) {
	if err := k.validatePutIngressApplicationGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIngressApplicationGateway",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutKeyVaultSecretsProvider(value *KubernetesClusterKeyVaultSecretsProvider) {
	if err := k.validatePutKeyVaultSecretsProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKeyVaultSecretsProvider",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutKubeletIdentity(value *KubernetesClusterKubeletIdentity) {
	if err := k.validatePutKubeletIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKubeletIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutLinuxProfile(value *KubernetesClusterLinuxProfile) {
	if err := k.validatePutLinuxProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLinuxProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutMaintenanceWindow(value *KubernetesClusterMaintenanceWindow) {
	if err := k.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutNetworkProfile(value *KubernetesClusterNetworkProfile) {
	if err := k.validatePutNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutOmsAgent(value *KubernetesClusterOmsAgent) {
	if err := k.validatePutOmsAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOmsAgent",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutRoleBasedAccessControl(value *KubernetesClusterRoleBasedAccessControl) {
	if err := k.validatePutRoleBasedAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRoleBasedAccessControl",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutServicePrincipal(value *KubernetesClusterServicePrincipal) {
	if err := k.validatePutServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServicePrincipal",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutTimeouts(value *KubernetesClusterTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) PutWindowsProfile(value *KubernetesClusterWindowsProfile) {
	if err := k.validatePutWindowsProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWindowsProfile",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAciConnectorLinux() {
	_jsii_.InvokeVoid(
		k,
		"resetAciConnectorLinux",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAddonProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetAddonProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetApiServerAuthorizedIpRanges() {
	_jsii_.InvokeVoid(
		k,
		"resetApiServerAuthorizedIpRanges",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAutomaticChannelUpgrade() {
	_jsii_.InvokeVoid(
		k,
		"resetAutomaticChannelUpgrade",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAutoScalerProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoScalerProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAzureActiveDirectoryRoleBasedAccessControl() {
	_jsii_.InvokeVoid(
		k,
		"resetAzureActiveDirectoryRoleBasedAccessControl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetAzurePolicyEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetAzurePolicyEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		k,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDnsPrefix() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsPrefix",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetDnsPrefixPrivateCluster() {
	_jsii_.InvokeVoid(
		k,
		"resetDnsPrefixPrivateCluster",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetEnablePodSecurityPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetEnablePodSecurityPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetHttpApplicationRoutingEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpApplicationRoutingEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetHttpProxyConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpProxyConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetIngressApplicationGateway() {
	_jsii_.InvokeVoid(
		k,
		"resetIngressApplicationGateway",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKeyVaultSecretsProvider() {
	_jsii_.InvokeVoid(
		k,
		"resetKeyVaultSecretsProvider",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKubeletIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetKubernetesVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetLinuxProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetLinuxProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetLocalAccountDisabled() {
	_jsii_.InvokeVoid(
		k,
		"resetLocalAccountDisabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		k,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNetworkProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetNodeResourceGroup() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeResourceGroup",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOmsAgent() {
	_jsii_.InvokeVoid(
		k,
		"resetOmsAgent",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOpenServiceMeshEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetOpenServiceMeshEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateClusterEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateClusterEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateClusterPublicFqdnEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateClusterPublicFqdnEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateDnsZoneId() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateDnsZoneId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPrivateLinkEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateLinkEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetRoleBasedAccessControl() {
	_jsii_.InvokeVoid(
		k,
		"resetRoleBasedAccessControl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetRoleBasedAccessControlEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetRoleBasedAccessControlEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetServicePrincipal() {
	_jsii_.InvokeVoid(
		k,
		"resetServicePrincipal",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetSkuTier() {
	_jsii_.InvokeVoid(
		k,
		"resetSkuTier",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) ResetWindowsProfile() {
	_jsii_.InvokeVoid(
		k,
		"resetWindowsProfile",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

