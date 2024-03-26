package mssqldatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/mssqldatabase/internal"
)

type MssqlDatabaseExtendedAuditingPolicyOutputReference interface {
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
	LogMonitoringEnabled() interface{}
	SetLogMonitoringEnabled(val interface{})
	LogMonitoringEnabledInput() interface{}
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	StorageAccountAccessKey() *string
	SetStorageAccountAccessKey(val *string)
	StorageAccountAccessKeyInput() *string
	StorageAccountAccessKeyIsSecondary() interface{}
	SetStorageAccountAccessKeyIsSecondary(val interface{})
	StorageAccountAccessKeyIsSecondaryInput() interface{}
	StorageAccountSubscriptionId() *string
	SetStorageAccountSubscriptionId(val *string)
	StorageAccountSubscriptionIdInput() *string
	StorageEndpoint() *string
	SetStorageEndpoint(val *string)
	StorageEndpointInput() *string
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
	ResetLogMonitoringEnabled()
	ResetRetentionInDays()
	ResetStorageAccountAccessKey()
	ResetStorageAccountAccessKeyIsSecondary()
	ResetStorageAccountSubscriptionId()
	ResetStorageEndpoint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlDatabaseExtendedAuditingPolicyOutputReference
type jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) LogMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) LogMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountAccessKeyIsSecondary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountAccessKeyIsSecondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountAccessKeyIsSecondaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountAccessKeyIsSecondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageAccountSubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSubscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) StorageEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlDatabaseExtendedAuditingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MssqlDatabaseExtendedAuditingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlDatabaseExtendedAuditingPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference{}

	_jsii_.Create(
		"azurerm.mssqlDatabase.MssqlDatabaseExtendedAuditingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMssqlDatabaseExtendedAuditingPolicyOutputReference_Override(m MssqlDatabaseExtendedAuditingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.mssqlDatabase.MssqlDatabaseExtendedAuditingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetLogMonitoringEnabled(val interface{}) {
	if err := j.validateSetLogMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetStorageAccountAccessKey(val *string) {
	if err := j.validateSetStorageAccountAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKey",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetStorageAccountAccessKeyIsSecondary(val interface{}) {
	if err := j.validateSetStorageAccountAccessKeyIsSecondaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountAccessKeyIsSecondary",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetStorageAccountSubscriptionId(val *string) {
	if err := j.validateSetStorageAccountSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountSubscriptionId",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetStorageEndpoint(val *string) {
	if err := j.validateSetStorageEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEndpoint",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetLogMonitoringEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetLogMonitoringEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		m,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetStorageAccountAccessKey() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountAccessKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetStorageAccountAccessKeyIsSecondary() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountAccessKeyIsSecondary",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetStorageAccountSubscriptionId() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageAccountSubscriptionId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ResetStorageEndpoint() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageEndpoint",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlDatabaseExtendedAuditingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

