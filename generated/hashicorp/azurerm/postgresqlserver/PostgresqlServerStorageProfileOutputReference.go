package postgresqlserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/postgresqlserver/internal"
)

type PostgresqlServerStorageProfileOutputReference interface {
	cdktf.ComplexObject
	AutoGrow() *string
	SetAutoGrow(val *string)
	AutoGrowInput() *string
	BackupRetentionDays() *float64
	SetBackupRetentionDays(val *float64)
	BackupRetentionDaysInput() *float64
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
	GeoRedundantBackup() *string
	SetGeoRedundantBackup(val *string)
	GeoRedundantBackupInput() *string
	InternalValue() *PostgresqlServerStorageProfile
	SetInternalValue(val *PostgresqlServerStorageProfile)
	StorageMb() *float64
	SetStorageMb(val *float64)
	StorageMbInput() *float64
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
	ResetAutoGrow()
	ResetBackupRetentionDays()
	ResetGeoRedundantBackup()
	ResetStorageMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PostgresqlServerStorageProfileOutputReference
type jsiiProxy_PostgresqlServerStorageProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) AutoGrow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoGrow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) AutoGrowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoGrowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) BackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) BackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GeoRedundantBackup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoRedundantBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GeoRedundantBackupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"geoRedundantBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) InternalValue() *PostgresqlServerStorageProfile {
	var returns *PostgresqlServerStorageProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) StorageMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) StorageMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPostgresqlServerStorageProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PostgresqlServerStorageProfileOutputReference {
	_init_.Initialize()

	if err := validateNewPostgresqlServerStorageProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostgresqlServerStorageProfileOutputReference{}

	_jsii_.Create(
		"azurerm.postgresqlServer.PostgresqlServerStorageProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPostgresqlServerStorageProfileOutputReference_Override(p PostgresqlServerStorageProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.postgresqlServer.PostgresqlServerStorageProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetAutoGrow(val *string) {
	if err := j.validateSetAutoGrowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoGrow",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetBackupRetentionDays(val *float64) {
	if err := j.validateSetBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetGeoRedundantBackup(val *string) {
	if err := j.validateSetGeoRedundantBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRedundantBackup",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetInternalValue(val *PostgresqlServerStorageProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetStorageMb(val *float64) {
	if err := j.validateSetStorageMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMb",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PostgresqlServerStorageProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ResetAutoGrow() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoGrow",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ResetBackupRetentionDays() {
	_jsii_.InvokeVoid(
		p,
		"resetBackupRetentionDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ResetGeoRedundantBackup() {
	_jsii_.InvokeVoid(
		p,
		"resetGeoRedundantBackup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ResetStorageMb() {
	_jsii_.InvokeVoid(
		p,
		"resetStorageMb",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostgresqlServerStorageProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

