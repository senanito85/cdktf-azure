package storagemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/storagemanagementpolicy/internal"
)

type StorageManagementPolicyRuleActionsOutputReference interface {
	cdktf.ComplexObject
	BaseBlob() StorageManagementPolicyRuleActionsBaseBlobOutputReference
	BaseBlobInput() *StorageManagementPolicyRuleActionsBaseBlob
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
	InternalValue() *StorageManagementPolicyRuleActions
	SetInternalValue(val *StorageManagementPolicyRuleActions)
	Snapshot() StorageManagementPolicyRuleActionsSnapshotOutputReference
	SnapshotInput() *StorageManagementPolicyRuleActionsSnapshot
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() StorageManagementPolicyRuleActionsVersionOutputReference
	VersionInput() *StorageManagementPolicyRuleActionsVersion
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
	PutBaseBlob(value *StorageManagementPolicyRuleActionsBaseBlob)
	PutSnapshot(value *StorageManagementPolicyRuleActionsSnapshot)
	PutVersion(value *StorageManagementPolicyRuleActionsVersion)
	ResetBaseBlob()
	ResetSnapshot()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageManagementPolicyRuleActionsOutputReference
type jsiiProxy_StorageManagementPolicyRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) BaseBlob() StorageManagementPolicyRuleActionsBaseBlobOutputReference {
	var returns StorageManagementPolicyRuleActionsBaseBlobOutputReference
	_jsii_.Get(
		j,
		"baseBlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) BaseBlobInput() *StorageManagementPolicyRuleActionsBaseBlob {
	var returns *StorageManagementPolicyRuleActionsBaseBlob
	_jsii_.Get(
		j,
		"baseBlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InternalValue() *StorageManagementPolicyRuleActions {
	var returns *StorageManagementPolicyRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Snapshot() StorageManagementPolicyRuleActionsSnapshotOutputReference {
	var returns StorageManagementPolicyRuleActionsSnapshotOutputReference
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) SnapshotInput() *StorageManagementPolicyRuleActionsSnapshot {
	var returns *StorageManagementPolicyRuleActionsSnapshot
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Version() StorageManagementPolicyRuleActionsVersionOutputReference {
	var returns StorageManagementPolicyRuleActionsVersionOutputReference
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) VersionInput() *StorageManagementPolicyRuleActionsVersion {
	var returns *StorageManagementPolicyRuleActionsVersion
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewStorageManagementPolicyRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageManagementPolicyRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageManagementPolicyRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageManagementPolicyRuleActionsOutputReference{}

	_jsii_.Create(
		"azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageManagementPolicyRuleActionsOutputReference_Override(s StorageManagementPolicyRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.storageManagementPolicy.StorageManagementPolicyRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference)SetInternalValue(val *StorageManagementPolicyRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutBaseBlob(value *StorageManagementPolicyRuleActionsBaseBlob) {
	if err := s.validatePutBaseBlobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBaseBlob",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutSnapshot(value *StorageManagementPolicyRuleActionsSnapshot) {
	if err := s.validatePutSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSnapshot",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) PutVersion(value *StorageManagementPolicyRuleActionsVersion) {
	if err := s.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVersion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetBaseBlob() {
	_jsii_.InvokeVoid(
		s,
		"resetBaseBlob",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageManagementPolicyRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

