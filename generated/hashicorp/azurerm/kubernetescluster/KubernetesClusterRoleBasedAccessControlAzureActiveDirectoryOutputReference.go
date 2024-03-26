package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hashicorp/terraform-cdk/examples/go/azure/generated/hashicorp/azurerm/kubernetescluster/internal"
)

type KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	AdminGroupObjectIds() *[]*string
	SetAdminGroupObjectIds(val *[]*string)
	AdminGroupObjectIdsInput() *[]*string
	AzureRbacEnabled() interface{}
	SetAzureRbacEnabled(val interface{})
	AzureRbacEnabledInput() interface{}
	ClientAppId() *string
	SetClientAppId(val *string)
	ClientAppIdInput() *string
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
	InternalValue() *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory
	SetInternalValue(val *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory)
	Managed() interface{}
	SetManaged(val interface{})
	ManagedInput() interface{}
	ServerAppId() *string
	SetServerAppId(val *string)
	ServerAppIdInput() *string
	ServerAppSecret() *string
	SetServerAppSecret(val *string)
	ServerAppSecretInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetAdminGroupObjectIds()
	ResetAzureRbacEnabled()
	ResetClientAppId()
	ResetManaged()
	ResetServerAppId()
	ResetServerAppSecret()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference
type jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) AdminGroupObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupObjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) AdminGroupObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminGroupObjectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) AzureRbacEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureRbacEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) AzureRbacEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureRbacEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ClientAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ClientAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) InternalValue() *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory {
	var returns *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) Managed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ManagedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ServerAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ServerAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ServerAppSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ServerAppSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverAppSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference{}

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference_Override(k KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"azurerm.kubernetesCluster.KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetAdminGroupObjectIds(val *[]*string) {
	if err := j.validateSetAdminGroupObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminGroupObjectIds",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetAzureRbacEnabled(val interface{}) {
	if err := j.validateSetAzureRbacEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureRbacEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetClientAppId(val *string) {
	if err := j.validateSetClientAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAppId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetInternalValue(val *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetManaged(val interface{}) {
	if err := j.validateSetManagedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managed",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetServerAppId(val *string) {
	if err := j.validateSetServerAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAppId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetServerAppSecret(val *string) {
	if err := j.validateSetServerAppSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverAppSecret",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetAdminGroupObjectIds() {
	_jsii_.InvokeVoid(
		k,
		"resetAdminGroupObjectIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetAzureRbacEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetAzureRbacEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetClientAppId() {
	_jsii_.InvokeVoid(
		k,
		"resetClientAppId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetManaged() {
	_jsii_.InvokeVoid(
		k,
		"resetManaged",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetServerAppId() {
	_jsii_.InvokeVoid(
		k,
		"resetServerAppId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetServerAppSecret() {
	_jsii_.InvokeVoid(
		k,
		"resetServerAppSecret",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		k,
		"resetTenantId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterRoleBasedAccessControlAzureActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

