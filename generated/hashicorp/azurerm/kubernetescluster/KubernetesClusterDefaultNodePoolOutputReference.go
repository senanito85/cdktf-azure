package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/kubernetescluster/internal"
)

type KubernetesClusterDefaultNodePoolOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableAutoScaling() interface{}
	SetEnableAutoScaling(val interface{})
	EnableAutoScalingInput() interface{}
	EnableHostEncryption() interface{}
	SetEnableHostEncryption(val interface{})
	EnableHostEncryptionInput() interface{}
	EnableNodePublicIp() interface{}
	SetEnableNodePublicIp(val interface{})
	EnableNodePublicIpInput() interface{}
	FipsEnabled() interface{}
	SetFipsEnabled(val interface{})
	FipsEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterDefaultNodePool
	SetInternalValue(val *KubernetesClusterDefaultNodePool)
	KubeletConfig() KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
	KubeletConfigInput() *KubernetesClusterDefaultNodePoolKubeletConfig
	KubeletDiskType() *string
	SetKubeletDiskType(val *string)
	KubeletDiskTypeInput() *string
	LinuxOsConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
	LinuxOsConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfig
	MaxCount() *float64
	SetMaxCount(val *float64)
	MaxCountInput() *float64
	MaxPods() *float64
	SetMaxPods(val *float64)
	MaxPodsInput() *float64
	MinCount() *float64
	SetMinCount(val *float64)
	MinCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLabels() *map[string]*string
	SetNodeLabels(val *map[string]*string)
	NodeLabelsInput() *map[string]*string
	NodePublicIpPrefixId() *string
	SetNodePublicIpPrefixId(val *string)
	NodePublicIpPrefixIdInput() *string
	NodeTaints() *[]*string
	SetNodeTaints(val *[]*string)
	NodeTaintsInput() *[]*string
	OnlyCriticalAddonsEnabled() interface{}
	SetOnlyCriticalAddonsEnabled(val interface{})
	OnlyCriticalAddonsEnabledInput() interface{}
	OrchestratorVersion() *string
	SetOrchestratorVersion(val *string)
	OrchestratorVersionInput() *string
	OsDiskSizeGb() *float64
	SetOsDiskSizeGb(val *float64)
	OsDiskSizeGbInput() *float64
	OsDiskType() *string
	SetOsDiskType(val *string)
	OsDiskTypeInput() *string
	OsSku() *string
	SetOsSku(val *string)
	OsSkuInput() *string
	PodSubnetId() *string
	SetPodSubnetId(val *string)
	PodSubnetIdInput() *string
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UltraSsdEnabled() interface{}
	SetUltraSsdEnabled(val interface{})
	UltraSsdEnabledInput() interface{}
	UpgradeSettings() KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference
	UpgradeSettingsInput() *KubernetesClusterDefaultNodePoolUpgradeSettings
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
	VnetSubnetId() *string
	SetVnetSubnetId(val *string)
	VnetSubnetIdInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutKubeletConfig(value *KubernetesClusterDefaultNodePoolKubeletConfig)
	PutLinuxOsConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfig)
	PutUpgradeSettings(value *KubernetesClusterDefaultNodePoolUpgradeSettings)
	ResetAvailabilityZones()
	ResetEnableAutoScaling()
	ResetEnableHostEncryption()
	ResetEnableNodePublicIp()
	ResetFipsEnabled()
	ResetKubeletConfig()
	ResetKubeletDiskType()
	ResetLinuxOsConfig()
	ResetMaxCount()
	ResetMaxPods()
	ResetMinCount()
	ResetNodeCount()
	ResetNodeLabels()
	ResetNodePublicIpPrefixId()
	ResetNodeTaints()
	ResetOnlyCriticalAddonsEnabled()
	ResetOrchestratorVersion()
	ResetOsDiskSizeGb()
	ResetOsDiskType()
	ResetOsSku()
	ResetPodSubnetId()
	ResetProximityPlacementGroupId()
	ResetTags()
	ResetType()
	ResetUltraSsdEnabled()
	ResetUpgradeSettings()
	ResetVnetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableAutoScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableAutoScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableHostEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableHostEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHostEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableNodePublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) EnableNodePublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNodePublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) FipsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) FipsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fipsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InternalValue() *KubernetesClusterDefaultNodePool {
	var returns *KubernetesClusterDefaultNodePool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletConfig() KubernetesClusterDefaultNodePoolKubeletConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletConfigInput() *KubernetesClusterDefaultNodePoolKubeletConfig {
	var returns *KubernetesClusterDefaultNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) KubeletDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeletDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) LinuxOsConfig() KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference {
	var returns KubernetesClusterDefaultNodePoolLinuxOsConfigOutputReference
	_jsii_.Get(
		j,
		"linuxOsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) LinuxOsConfigInput() *KubernetesClusterDefaultNodePoolLinuxOsConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfig
	_jsii_.Get(
		j,
		"linuxOsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxPods() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MaxPodsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MinCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) MinCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nodeLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodePublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodePublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodePublicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeTaints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) NodeTaintsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeTaintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OnlyCriticalAddonsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyCriticalAddonsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OnlyCriticalAddonsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyCriticalAddonsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OrchestratorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OrchestratorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orchestratorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"osDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) OsSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PodSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PodSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UltraSsdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UltraSsdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UpgradeSettings() KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference {
	var returns KubernetesClusterDefaultNodePoolUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) UpgradeSettingsInput() *KubernetesClusterDefaultNodePoolUpgradeSettings {
	var returns *KubernetesClusterDefaultNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VnetSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) VnetSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetSubnetIdInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterDefaultNodePoolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference{}

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolOutputReference_Override(k KubernetesClusterDefaultNodePoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetEnableAutoScaling(val interface{}) {
	if err := j.validateSetEnableAutoScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoScaling",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetEnableHostEncryption(val interface{}) {
	if err := j.validateSetEnableHostEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHostEncryption",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetEnableNodePublicIp(val interface{}) {
	if err := j.validateSetEnableNodePublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNodePublicIp",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetFipsEnabled(val interface{}) {
	if err := j.validateSetFipsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fipsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetInternalValue(val *KubernetesClusterDefaultNodePool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetKubeletDiskType(val *string) {
	if err := j.validateSetKubeletDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubeletDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetMaxCount(val *float64) {
	if err := j.validateSetMaxCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetMaxPods(val *float64) {
	if err := j.validateSetMaxPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetMinCount(val *float64) {
	if err := j.validateSetMinCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetNodeLabels(val *map[string]*string) {
	if err := j.validateSetNodeLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetNodePublicIpPrefixId(val *string) {
	if err := j.validateSetNodePublicIpPrefixIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodePublicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetNodeTaints(val *[]*string) {
	if err := j.validateSetNodeTaintsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTaints",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetOnlyCriticalAddonsEnabled(val interface{}) {
	if err := j.validateSetOnlyCriticalAddonsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyCriticalAddonsEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetOrchestratorVersion(val *string) {
	if err := j.validateSetOrchestratorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orchestratorVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetOsDiskSizeGb(val *float64) {
	if err := j.validateSetOsDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetOsDiskType(val *string) {
	if err := j.validateSetOsDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetOsSku(val *string) {
	if err := j.validateSetOsSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osSku",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetPodSubnetId(val *string) {
	if err := j.validateSetPodSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podSubnetId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetUltraSsdEnabled(val interface{}) {
	if err := j.validateSetUltraSsdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference)SetVnetSubnetId(val *string) {
	if err := j.validateSetVnetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetSubnetId",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutKubeletConfig(value *KubernetesClusterDefaultNodePoolKubeletConfig) {
	if err := k.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutLinuxOsConfig(value *KubernetesClusterDefaultNodePoolLinuxOsConfig) {
	if err := k.validatePutLinuxOsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putLinuxOsConfig",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) PutUpgradeSettings(value *KubernetesClusterDefaultNodePoolUpgradeSettings) {
	if err := k.validatePutUpgradeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putUpgradeSettings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		k,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableAutoScaling() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableAutoScaling",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableHostEncryption() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableHostEncryption",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetEnableNodePublicIp() {
	_jsii_.InvokeVoid(
		k,
		"resetEnableNodePublicIp",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetFipsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetFipsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetKubeletDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeletDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetLinuxOsConfig() {
	_jsii_.InvokeVoid(
		k,
		"resetLinuxOsConfig",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMaxCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMaxPods() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxPods",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetMinCount() {
	_jsii_.InvokeVoid(
		k,
		"resetMinCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodePublicIpPrefixId() {
	_jsii_.InvokeVoid(
		k,
		"resetNodePublicIpPrefixId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetNodeTaints() {
	_jsii_.InvokeVoid(
		k,
		"resetNodeTaints",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOnlyCriticalAddonsEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetOnlyCriticalAddonsEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOrchestratorVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetOrchestratorVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsDiskSizeGb() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskSizeGb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsDiskType() {
	_jsii_.InvokeVoid(
		k,
		"resetOsDiskType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetOsSku() {
	_jsii_.InvokeVoid(
		k,
		"resetOsSku",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetPodSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetPodSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		k,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		k,
		"resetType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetUltraSsdEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetUltraSsdEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		k,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ResetVnetSubnetId() {
	_jsii_.InvokeVoid(
		k,
		"resetVnetSubnetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

