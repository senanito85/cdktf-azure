package frontdoorfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/frontdoorfirewallpolicy/internal"
)

type FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference interface {
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
	MatchVariable() *string
	SetMatchVariable(val *string)
	MatchVariableInput() *string
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Selector() *string
	SetSelector(val *string)
	SelectorInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference
type jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) MatchVariable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) MatchVariableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) Selector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) SelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference {
	_init_.Initialize()

	if err := validateNewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference{}

	_jsii_.Create(
		"azurerm.frontdoorFirewallPolicy.FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference_Override(f FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.frontdoorFirewallPolicy.FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetMatchVariable(val *string) {
	if err := j.validateSetMatchVariableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchVariable",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetSelector(val *string) {
	if err := j.validateSetSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selector",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorFirewallPolicyManagedRuleOverrideRuleExclusionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

