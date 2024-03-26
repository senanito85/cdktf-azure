package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/eventgrideventsubscription/internal"
)

type EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	_init_.Initialize()

	if err := validateNewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList{}

	_jsii_.Create(
		"azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList_Override(e EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := e.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		e,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

