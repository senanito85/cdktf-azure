package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/kubernetescluster/internal"
)

type KubernetesClusterAddonProfileOutputReference interface {
	cdktf.ComplexObject
	AciConnectorLinux() KubernetesClusterAddonProfileAciConnectorLinuxOutputReference
	AciConnectorLinuxInput() *KubernetesClusterAddonProfileAciConnectorLinux
	AzureKeyvaultSecretsProvider() KubernetesClusterAddonProfileAzureKeyvaultSecretsProviderOutputReference
	AzureKeyvaultSecretsProviderInput() *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider
	AzurePolicy() KubernetesClusterAddonProfileAzurePolicyOutputReference
	AzurePolicyInput() *KubernetesClusterAddonProfileAzurePolicy
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
	// Experimental.
	Fqn() *string
	HttpApplicationRouting() KubernetesClusterAddonProfileHttpApplicationRoutingOutputReference
	HttpApplicationRoutingInput() *KubernetesClusterAddonProfileHttpApplicationRouting
	IngressApplicationGateway() KubernetesClusterAddonProfileIngressApplicationGatewayOutputReference
	IngressApplicationGatewayInput() *KubernetesClusterAddonProfileIngressApplicationGateway
	InternalValue() *KubernetesClusterAddonProfile
	SetInternalValue(val *KubernetesClusterAddonProfile)
	KubeDashboard() KubernetesClusterAddonProfileKubeDashboardOutputReference
	KubeDashboardInput() *KubernetesClusterAddonProfileKubeDashboard
	OmsAgent() KubernetesClusterAddonProfileOmsAgentOutputReference
	OmsAgentInput() *KubernetesClusterAddonProfileOmsAgent
	OpenServiceMesh() KubernetesClusterAddonProfileOpenServiceMeshOutputReference
	OpenServiceMeshInput() *KubernetesClusterAddonProfileOpenServiceMesh
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAciConnectorLinux(value *KubernetesClusterAddonProfileAciConnectorLinux)
	PutAzureKeyvaultSecretsProvider(value *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider)
	PutAzurePolicy(value *KubernetesClusterAddonProfileAzurePolicy)
	PutHttpApplicationRouting(value *KubernetesClusterAddonProfileHttpApplicationRouting)
	PutIngressApplicationGateway(value *KubernetesClusterAddonProfileIngressApplicationGateway)
	PutKubeDashboard(value *KubernetesClusterAddonProfileKubeDashboard)
	PutOmsAgent(value *KubernetesClusterAddonProfileOmsAgent)
	PutOpenServiceMesh(value *KubernetesClusterAddonProfileOpenServiceMesh)
	ResetAciConnectorLinux()
	ResetAzureKeyvaultSecretsProvider()
	ResetAzurePolicy()
	ResetHttpApplicationRouting()
	ResetIngressApplicationGateway()
	ResetKubeDashboard()
	ResetOmsAgent()
	ResetOpenServiceMesh()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterAddonProfileOutputReference
type jsiiProxy_KubernetesClusterAddonProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AciConnectorLinux() KubernetesClusterAddonProfileAciConnectorLinuxOutputReference {
	var returns KubernetesClusterAddonProfileAciConnectorLinuxOutputReference
	_jsii_.Get(
		j,
		"aciConnectorLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AciConnectorLinuxInput() *KubernetesClusterAddonProfileAciConnectorLinux {
	var returns *KubernetesClusterAddonProfileAciConnectorLinux
	_jsii_.Get(
		j,
		"aciConnectorLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AzureKeyvaultSecretsProvider() KubernetesClusterAddonProfileAzureKeyvaultSecretsProviderOutputReference {
	var returns KubernetesClusterAddonProfileAzureKeyvaultSecretsProviderOutputReference
	_jsii_.Get(
		j,
		"azureKeyvaultSecretsProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AzureKeyvaultSecretsProviderInput() *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider {
	var returns *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider
	_jsii_.Get(
		j,
		"azureKeyvaultSecretsProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AzurePolicy() KubernetesClusterAddonProfileAzurePolicyOutputReference {
	var returns KubernetesClusterAddonProfileAzurePolicyOutputReference
	_jsii_.Get(
		j,
		"azurePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) AzurePolicyInput() *KubernetesClusterAddonProfileAzurePolicy {
	var returns *KubernetesClusterAddonProfileAzurePolicy
	_jsii_.Get(
		j,
		"azurePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) HttpApplicationRouting() KubernetesClusterAddonProfileHttpApplicationRoutingOutputReference {
	var returns KubernetesClusterAddonProfileHttpApplicationRoutingOutputReference
	_jsii_.Get(
		j,
		"httpApplicationRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) HttpApplicationRoutingInput() *KubernetesClusterAddonProfileHttpApplicationRouting {
	var returns *KubernetesClusterAddonProfileHttpApplicationRouting
	_jsii_.Get(
		j,
		"httpApplicationRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) IngressApplicationGateway() KubernetesClusterAddonProfileIngressApplicationGatewayOutputReference {
	var returns KubernetesClusterAddonProfileIngressApplicationGatewayOutputReference
	_jsii_.Get(
		j,
		"ingressApplicationGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) IngressApplicationGatewayInput() *KubernetesClusterAddonProfileIngressApplicationGateway {
	var returns *KubernetesClusterAddonProfileIngressApplicationGateway
	_jsii_.Get(
		j,
		"ingressApplicationGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) InternalValue() *KubernetesClusterAddonProfile {
	var returns *KubernetesClusterAddonProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) KubeDashboard() KubernetesClusterAddonProfileKubeDashboardOutputReference {
	var returns KubernetesClusterAddonProfileKubeDashboardOutputReference
	_jsii_.Get(
		j,
		"kubeDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) KubeDashboardInput() *KubernetesClusterAddonProfileKubeDashboard {
	var returns *KubernetesClusterAddonProfileKubeDashboard
	_jsii_.Get(
		j,
		"kubeDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) OmsAgent() KubernetesClusterAddonProfileOmsAgentOutputReference {
	var returns KubernetesClusterAddonProfileOmsAgentOutputReference
	_jsii_.Get(
		j,
		"omsAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) OmsAgentInput() *KubernetesClusterAddonProfileOmsAgent {
	var returns *KubernetesClusterAddonProfileOmsAgent
	_jsii_.Get(
		j,
		"omsAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) OpenServiceMesh() KubernetesClusterAddonProfileOpenServiceMeshOutputReference {
	var returns KubernetesClusterAddonProfileOpenServiceMeshOutputReference
	_jsii_.Get(
		j,
		"openServiceMesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) OpenServiceMeshInput() *KubernetesClusterAddonProfileOpenServiceMesh {
	var returns *KubernetesClusterAddonProfileOpenServiceMesh
	_jsii_.Get(
		j,
		"openServiceMeshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterAddonProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterAddonProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterAddonProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterAddonProfileOutputReference{}

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterAddonProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterAddonProfileOutputReference_Override(k KubernetesClusterAddonProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterAddonProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference)SetInternalValue(val *KubernetesClusterAddonProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAddonProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutAciConnectorLinux(value *KubernetesClusterAddonProfileAciConnectorLinux) {
	if err := k.validatePutAciConnectorLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAciConnectorLinux",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutAzureKeyvaultSecretsProvider(value *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider) {
	if err := k.validatePutAzureKeyvaultSecretsProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAzureKeyvaultSecretsProvider",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutAzurePolicy(value *KubernetesClusterAddonProfileAzurePolicy) {
	if err := k.validatePutAzurePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAzurePolicy",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutHttpApplicationRouting(value *KubernetesClusterAddonProfileHttpApplicationRouting) {
	if err := k.validatePutHttpApplicationRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putHttpApplicationRouting",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutIngressApplicationGateway(value *KubernetesClusterAddonProfileIngressApplicationGateway) {
	if err := k.validatePutIngressApplicationGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIngressApplicationGateway",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutKubeDashboard(value *KubernetesClusterAddonProfileKubeDashboard) {
	if err := k.validatePutKubeDashboardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKubeDashboard",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutOmsAgent(value *KubernetesClusterAddonProfileOmsAgent) {
	if err := k.validatePutOmsAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOmsAgent",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) PutOpenServiceMesh(value *KubernetesClusterAddonProfileOpenServiceMesh) {
	if err := k.validatePutOpenServiceMeshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOpenServiceMesh",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetAciConnectorLinux() {
	_jsii_.InvokeVoid(
		k,
		"resetAciConnectorLinux",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetAzureKeyvaultSecretsProvider() {
	_jsii_.InvokeVoid(
		k,
		"resetAzureKeyvaultSecretsProvider",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetAzurePolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetAzurePolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetHttpApplicationRouting() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpApplicationRouting",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetIngressApplicationGateway() {
	_jsii_.InvokeVoid(
		k,
		"resetIngressApplicationGateway",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetKubeDashboard() {
	_jsii_.InvokeVoid(
		k,
		"resetKubeDashboard",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetOmsAgent() {
	_jsii_.InvokeVoid(
		k,
		"resetOmsAgent",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ResetOpenServiceMesh() {
	_jsii_.InvokeVoid(
		k,
		"resetOpenServiceMesh",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterAddonProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

