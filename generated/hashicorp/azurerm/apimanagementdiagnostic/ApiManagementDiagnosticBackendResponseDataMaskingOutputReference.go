package apimanagementdiagnostic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/apimanagementdiagnostic/internal"
)

type ApiManagementDiagnosticBackendResponseDataMaskingOutputReference interface {
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
	Headers() ApiManagementDiagnosticBackendResponseDataMaskingHeadersList
	HeadersInput() interface{}
	InternalValue() *ApiManagementDiagnosticBackendResponseDataMasking
	SetInternalValue(val *ApiManagementDiagnosticBackendResponseDataMasking)
	QueryParams() ApiManagementDiagnosticBackendResponseDataMaskingQueryParamsList
	QueryParamsInput() interface{}
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
	PutHeaders(value interface{})
	PutQueryParams(value interface{})
	ResetHeaders()
	ResetQueryParams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementDiagnosticBackendResponseDataMaskingOutputReference
type jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) Headers() ApiManagementDiagnosticBackendResponseDataMaskingHeadersList {
	var returns ApiManagementDiagnosticBackendResponseDataMaskingHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) InternalValue() *ApiManagementDiagnosticBackendResponseDataMasking {
	var returns *ApiManagementDiagnosticBackendResponseDataMasking
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) QueryParams() ApiManagementDiagnosticBackendResponseDataMaskingQueryParamsList {
	var returns ApiManagementDiagnosticBackendResponseDataMaskingQueryParamsList
	_jsii_.Get(
		j,
		"queryParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) QueryParamsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementDiagnosticBackendResponseDataMaskingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementDiagnosticBackendResponseDataMaskingOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementDiagnosticBackendResponseDataMaskingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference{}

	_jsii_.Create(
		"azurerm.apiManagementDiagnostic.ApiManagementDiagnosticBackendResponseDataMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementDiagnosticBackendResponseDataMaskingOutputReference_Override(a ApiManagementDiagnosticBackendResponseDataMaskingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.apiManagementDiagnostic.ApiManagementDiagnosticBackendResponseDataMaskingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference)SetInternalValue(val *ApiManagementDiagnosticBackendResponseDataMasking) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) PutHeaders(value interface{}) {
	if err := a.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) PutQueryParams(value interface{}) {
	if err := a.validatePutQueryParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryParams",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ResetQueryParams() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementDiagnosticBackendResponseDataMaskingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

