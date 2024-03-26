package networkconnectionmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/networkconnectionmonitor/internal"
)

type NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference interface {
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
	DestinationPortBehavior() *string
	SetDestinationPortBehavior(val *string)
	DestinationPortBehaviorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkConnectionMonitorTestConfigurationTcpConfiguration
	SetInternalValue(val *NetworkConnectionMonitorTestConfigurationTcpConfiguration)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TraceRouteEnabled() interface{}
	SetTraceRouteEnabled(val interface{})
	TraceRouteEnabledInput() interface{}
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
	ResetDestinationPortBehavior()
	ResetTraceRouteEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference
type jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) DestinationPortBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) DestinationPortBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) InternalValue() *NetworkConnectionMonitorTestConfigurationTcpConfiguration {
	var returns *NetworkConnectionMonitorTestConfigurationTcpConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) TraceRouteEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceRouteEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) TraceRouteEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"traceRouteEnabledInput",
		&returns,
	)
	return returns
}


func NewNetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference{}

	_jsii_.Create(
		"azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference_Override(n NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.networkConnectionMonitor.NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetDestinationPortBehavior(val *string) {
	if err := j.validateSetDestinationPortBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortBehavior",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetInternalValue(val *NetworkConnectionMonitorTestConfigurationTcpConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference)SetTraceRouteEnabled(val interface{}) {
	if err := j.validateSetTraceRouteEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"traceRouteEnabled",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ResetDestinationPortBehavior() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationPortBehavior",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ResetTraceRouteEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetTraceRouteEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectionMonitorTestConfigurationTcpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

