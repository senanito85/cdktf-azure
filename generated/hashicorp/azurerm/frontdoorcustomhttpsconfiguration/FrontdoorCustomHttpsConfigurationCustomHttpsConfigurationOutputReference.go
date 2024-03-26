package frontdoorcustomhttpsconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/frontdoorcustomhttpsconfiguration/internal"
)

type FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference interface {
	cdktf.ComplexObject
	AzureKeyVaultCertificateSecretName() *string
	SetAzureKeyVaultCertificateSecretName(val *string)
	AzureKeyVaultCertificateSecretNameInput() *string
	AzureKeyVaultCertificateSecretVersion() *string
	SetAzureKeyVaultCertificateSecretVersion(val *string)
	AzureKeyVaultCertificateSecretVersionInput() *string
	AzureKeyVaultCertificateVaultId() *string
	SetAzureKeyVaultCertificateVaultId(val *string)
	AzureKeyVaultCertificateVaultIdInput() *string
	CertificateSource() *string
	SetCertificateSource(val *string)
	CertificateSourceInput() *string
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
	InternalValue() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	SetInternalValue(val *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration)
	MinimumTlsVersion() *string
	ProvisioningState() *string
	ProvisioningSubstate() *string
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
	ResetAzureKeyVaultCertificateSecretName()
	ResetAzureKeyVaultCertificateSecretVersion()
	ResetAzureKeyVaultCertificateVaultId()
	ResetCertificateSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference
type jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CertificateSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InternalValue() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration {
	var returns *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ProvisioningState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ProvisioningSubstate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningSubstate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference{}

	_jsii_.Create(
		"azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference_Override(f FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetAzureKeyVaultCertificateSecretName(val *string) {
	if err := j.validateSetAzureKeyVaultCertificateSecretNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateSecretName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetAzureKeyVaultCertificateSecretVersion(val *string) {
	if err := j.validateSetAzureKeyVaultCertificateSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateSecretVersion",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetAzureKeyVaultCertificateVaultId(val *string) {
	if err := j.validateSetAzureKeyVaultCertificateVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateVaultId",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetCertificateSource(val *string) {
	if err := j.validateSetCertificateSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateSource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetInternalValue(val *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateSecretName() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateSecretName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateSecretVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateSecretVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateVaultId() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateVaultId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetCertificateSource() {
	_jsii_.InvokeVoid(
		f,
		"resetCertificateSource",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

