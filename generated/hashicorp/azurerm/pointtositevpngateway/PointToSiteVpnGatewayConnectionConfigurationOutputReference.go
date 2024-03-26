package pointtositevpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/pointtositevpngateway/internal"
)

type PointToSiteVpnGatewayConnectionConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *PointToSiteVpnGatewayConnectionConfiguration
	SetInternalValue(val *PointToSiteVpnGatewayConnectionConfiguration)
	InternetSecurityEnabled() interface{}
	SetInternetSecurityEnabled(val interface{})
	InternetSecurityEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Route() PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
	RouteInput() *PointToSiteVpnGatewayConnectionConfigurationRoute
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpnClientAddressPool() PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference
	VpnClientAddressPoolInput() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
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
	PutRoute(value *PointToSiteVpnGatewayConnectionConfigurationRoute)
	PutVpnClientAddressPool(value *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool)
	ResetInternetSecurityEnabled()
	ResetRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PointToSiteVpnGatewayConnectionConfigurationOutputReference
type jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternalValue() *PointToSiteVpnGatewayConnectionConfiguration {
	var returns *PointToSiteVpnGatewayConnectionConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternetSecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InternetSecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Route() PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationRouteOutputReference
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) RouteInput() *PointToSiteVpnGatewayConnectionConfigurationRoute {
	var returns *PointToSiteVpnGatewayConnectionConfigurationRoute
	_jsii_.Get(
		j,
		"routeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) VpnClientAddressPool() PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference {
	var returns PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPoolOutputReference
	_jsii_.Get(
		j,
		"vpnClientAddressPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) VpnClientAddressPoolInput() *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool {
	var returns *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool
	_jsii_.Get(
		j,
		"vpnClientAddressPoolInput",
		&returns,
	)
	return returns
}


func NewPointToSiteVpnGatewayConnectionConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PointToSiteVpnGatewayConnectionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPointToSiteVpnGatewayConnectionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference{}

	_jsii_.Create(
		"azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPointToSiteVpnGatewayConnectionConfigurationOutputReference_Override(p PointToSiteVpnGatewayConnectionConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.pointToSiteVpnGateway.PointToSiteVpnGatewayConnectionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetInternalValue(val *PointToSiteVpnGatewayConnectionConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetInternetSecurityEnabled(val interface{}) {
	if err := j.validateSetInternetSecurityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetSecurityEnabled",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) PutRoute(value *PointToSiteVpnGatewayConnectionConfigurationRoute) {
	if err := p.validatePutRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRoute",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) PutVpnClientAddressPool(value *PointToSiteVpnGatewayConnectionConfigurationVpnClientAddressPool) {
	if err := p.validatePutVpnClientAddressPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVpnClientAddressPool",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ResetInternetSecurityEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetInternetSecurityEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ResetRoute() {
	_jsii_.InvokeVoid(
		p,
		"resetRoute",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PointToSiteVpnGatewayConnectionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

