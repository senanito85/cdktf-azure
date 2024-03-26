package apimanagementapidiagnostic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/apimanagementapidiagnostic/internal"
)

type ApiManagementApiDiagnosticBackendRequestOutputReference interface {
	cdktf.ComplexObject
	BodyBytes() *float64
	SetBodyBytes(val *float64)
	BodyBytesInput() *float64
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
	DataMasking() ApiManagementApiDiagnosticBackendRequestDataMaskingOutputReference
	DataMaskingInput() *ApiManagementApiDiagnosticBackendRequestDataMasking
	// Experimental.
	Fqn() *string
	HeadersToLog() *[]*string
	SetHeadersToLog(val *[]*string)
	HeadersToLogInput() *[]*string
	InternalValue() *ApiManagementApiDiagnosticBackendRequest
	SetInternalValue(val *ApiManagementApiDiagnosticBackendRequest)
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
	PutDataMasking(value *ApiManagementApiDiagnosticBackendRequestDataMasking)
	ResetBodyBytes()
	ResetDataMasking()
	ResetHeadersToLog()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementApiDiagnosticBackendRequestOutputReference
type jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) BodyBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bodyBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) BodyBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bodyBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) DataMasking() ApiManagementApiDiagnosticBackendRequestDataMaskingOutputReference {
	var returns ApiManagementApiDiagnosticBackendRequestDataMaskingOutputReference
	_jsii_.Get(
		j,
		"dataMasking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) DataMaskingInput() *ApiManagementApiDiagnosticBackendRequestDataMasking {
	var returns *ApiManagementApiDiagnosticBackendRequestDataMasking
	_jsii_.Get(
		j,
		"dataMaskingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) HeadersToLog() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersToLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) HeadersToLogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersToLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) InternalValue() *ApiManagementApiDiagnosticBackendRequest {
	var returns *ApiManagementApiDiagnosticBackendRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementApiDiagnosticBackendRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementApiDiagnosticBackendRequestOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementApiDiagnosticBackendRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference{}

	_jsii_.Create(
		"azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnosticBackendRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementApiDiagnosticBackendRequestOutputReference_Override(a ApiManagementApiDiagnosticBackendRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.apiManagementApiDiagnostic.ApiManagementApiDiagnosticBackendRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetBodyBytes(val *float64) {
	if err := j.validateSetBodyBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bodyBytes",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetHeadersToLog(val *[]*string) {
	if err := j.validateSetHeadersToLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headersToLog",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetInternalValue(val *ApiManagementApiDiagnosticBackendRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) PutDataMasking(value *ApiManagementApiDiagnosticBackendRequestDataMasking) {
	if err := a.validatePutDataMaskingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataMasking",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ResetBodyBytes() {
	_jsii_.InvokeVoid(
		a,
		"resetBodyBytes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ResetDataMasking() {
	_jsii_.InvokeVoid(
		a,
		"resetDataMasking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ResetHeadersToLog() {
	_jsii_.InvokeVoid(
		a,
		"resetHeadersToLog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiManagementApiDiagnosticBackendRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

