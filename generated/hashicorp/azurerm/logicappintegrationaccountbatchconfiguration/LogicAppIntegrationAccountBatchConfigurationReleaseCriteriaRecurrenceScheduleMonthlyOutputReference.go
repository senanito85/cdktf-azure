package logicappintegrationaccountbatchconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/logicappintegrationaccountbatchconfiguration/internal"
)

type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Week() *float64
	SetWeek(val *float64)
	Weekday() *string
	SetWeekday(val *string)
	WeekdayInput() *string
	WeekInput() *float64
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

// The jsii proxy struct for LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference
type jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) Week() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"week",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) Weekday() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekday",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) WeekdayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weekdayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) WeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weekInput",
		&returns,
	)
	return returns
}


func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference{}

	_jsii_.Create(
		"azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference_Override(l LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.logicAppIntegrationAccountBatchConfiguration.LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetWeek(val *float64) {
	if err := j.validateSetWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"week",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference)SetWeekday(val *string) {
	if err := j.validateSetWeekdayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weekday",
		val,
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

