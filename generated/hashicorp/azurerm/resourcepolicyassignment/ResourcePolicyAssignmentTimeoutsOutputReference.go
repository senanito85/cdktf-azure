package resourcepolicyassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/resourcepolicyassignment/internal"
)

type ResourcePolicyAssignmentTimeoutsOutputReference interface {
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResourcePolicyAssignmentTimeoutsOutputReference
type jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewResourcePolicyAssignmentTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ResourcePolicyAssignmentTimeoutsOutputReference {
	_init_.Initialize()

	if err := validateNewResourcePolicyAssignmentTimeoutsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference{}

	_jsii_.Create(
		"azurerm.resourcePolicyAssignment.ResourcePolicyAssignmentTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResourcePolicyAssignmentTimeoutsOutputReference_Override(r ResourcePolicyAssignmentTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.resourcePolicyAssignment.ResourcePolicyAssignmentTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetCreate(val *string) {
	if err := j.validateSetCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetDelete(val *string) {
	if err := j.validateSetDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetRead(val *string) {
	if err := j.validateSetReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference)SetUpdate(val *string) {
	if err := j.validateSetUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		r,
		"resetCreate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		r,
		"resetDelete",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		r,
		"resetRead",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcePolicyAssignmentTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

