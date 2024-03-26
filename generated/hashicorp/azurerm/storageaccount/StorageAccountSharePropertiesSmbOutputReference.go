package storageaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/storageaccount/internal"
)

type StorageAccountSharePropertiesSmbOutputReference interface {
	cdktf.ComplexObject
	AuthenticationTypes() *[]*string
	SetAuthenticationTypes(val *[]*string)
	AuthenticationTypesInput() *[]*string
	ChannelEncryptionType() *[]*string
	SetChannelEncryptionType(val *[]*string)
	ChannelEncryptionTypeInput() *[]*string
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
	InternalValue() *StorageAccountSharePropertiesSmb
	SetInternalValue(val *StorageAccountSharePropertiesSmb)
	KerberosTicketEncryptionType() *[]*string
	SetKerberosTicketEncryptionType(val *[]*string)
	KerberosTicketEncryptionTypeInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Versions() *[]*string
	SetVersions(val *[]*string)
	VersionsInput() *[]*string
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
	ResetAuthenticationTypes()
	ResetChannelEncryptionType()
	ResetKerberosTicketEncryptionType()
	ResetVersions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageAccountSharePropertiesSmbOutputReference
type jsiiProxy_StorageAccountSharePropertiesSmbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) AuthenticationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) AuthenticationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authenticationTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ChannelEncryptionType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ChannelEncryptionTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InternalValue() *StorageAccountSharePropertiesSmb {
	var returns *StorageAccountSharePropertiesSmb
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) KerberosTicketEncryptionType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kerberosTicketEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) KerberosTicketEncryptionTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"kerberosTicketEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Versions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) VersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionsInput",
		&returns,
	)
	return returns
}


func NewStorageAccountSharePropertiesSmbOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageAccountSharePropertiesSmbOutputReference {
	_init_.Initialize()

	if err := validateNewStorageAccountSharePropertiesSmbOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageAccountSharePropertiesSmbOutputReference{}

	_jsii_.Create(
		"azurerm.storageAccount.StorageAccountSharePropertiesSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageAccountSharePropertiesSmbOutputReference_Override(s StorageAccountSharePropertiesSmbOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.storageAccount.StorageAccountSharePropertiesSmbOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetAuthenticationTypes(val *[]*string) {
	if err := j.validateSetAuthenticationTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationTypes",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetChannelEncryptionType(val *[]*string) {
	if err := j.validateSetChannelEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelEncryptionType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetInternalValue(val *StorageAccountSharePropertiesSmb) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetKerberosTicketEncryptionType(val *[]*string) {
	if err := j.validateSetKerberosTicketEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosTicketEncryptionType",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference)SetVersions(val *[]*string) {
	if err := j.validateSetVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versions",
		val,
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetAuthenticationTypes() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticationTypes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetChannelEncryptionType() {
	_jsii_.InvokeVoid(
		s,
		"resetChannelEncryptionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetKerberosTicketEncryptionType() {
	_jsii_.InvokeVoid(
		s,
		"resetKerberosTicketEncryptionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ResetVersions() {
	_jsii_.InvokeVoid(
		s,
		"resetVersions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageAccountSharePropertiesSmbOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

