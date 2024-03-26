package monitormetricalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/monitormetricalert/internal"
)

type MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference interface {
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
	ComponentId() *string
	SetComponentId(val *string)
	ComponentIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FailedLocationCount() *float64
	SetFailedLocationCount(val *float64)
	FailedLocationCountInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria
	SetInternalValue(val *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebTestId() *string
	SetWebTestId(val *string)
	WebTestIdInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference
type jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ComponentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ComponentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) FailedLocationCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedLocationCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) FailedLocationCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedLocationCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) InternalValue() *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria {
	var returns *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) WebTestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webTestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) WebTestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webTestIdInput",
		&returns,
	)
	return returns
}


func NewMonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference{}

	_jsii_.Create(
		"azurerm.monitorMetricAlert.MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference_Override(m MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.monitorMetricAlert.MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetComponentId(val *string) {
	if err := j.validateSetComponentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentId",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetFailedLocationCount(val *float64) {
	if err := j.validateSetFailedLocationCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failedLocationCount",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetInternalValue(val *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference)SetWebTestId(val *string) {
	if err := j.validateSetWebTestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webTestId",
		val,
	)
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

