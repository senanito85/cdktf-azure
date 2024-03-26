package databricksworkspace

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"azurerm.databricksWorkspace.DatabricksWorkspace",
		reflect.TypeOf((*DatabricksWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedKeyEnabled", GoGetter: "CustomerManagedKeyEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedKeyEnabledInput", GoGetter: "CustomerManagedKeyEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "customParameters", GoGetter: "CustomParameters"},
			_jsii_.MemberProperty{JsiiProperty: "customParametersInput", GoGetter: "CustomParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureEncryptionEnabled", GoGetter: "InfrastructureEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureEncryptionEnabledInput", GoGetter: "InfrastructureEncryptionEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerBackendAddressPoolId", GoGetter: "LoadBalancerBackendAddressPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerBackendAddressPoolIdInput", GoGetter: "LoadBalancerBackendAddressPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedResourceGroupId", GoGetter: "ManagedResourceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "managedResourceGroupName", GoGetter: "ManagedResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "managedResourceGroupNameInput", GoGetter: "ManagedResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedServicesCmkKeyVaultKeyId", GoGetter: "ManagedServicesCmkKeyVaultKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "managedServicesCmkKeyVaultKeyIdInput", GoGetter: "ManagedServicesCmkKeyVaultKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupRulesRequired", GoGetter: "NetworkSecurityGroupRulesRequired"},
			_jsii_.MemberProperty{JsiiProperty: "networkSecurityGroupRulesRequiredInput", GoGetter: "NetworkSecurityGroupRulesRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAccessEnabled", GoGetter: "PublicNetworkAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "publicNetworkAccessEnabledInput", GoGetter: "PublicNetworkAccessEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomParameters", GoMethod: "PutCustomParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerManagedKeyEnabled", GoMethod: "ResetCustomerManagedKeyEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomParameters", GoMethod: "ResetCustomParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInfrastructureEncryptionEnabled", GoMethod: "ResetInfrastructureEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancerBackendAddressPoolId", GoMethod: "ResetLoadBalancerBackendAddressPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedResourceGroupName", GoMethod: "ResetManagedResourceGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedServicesCmkKeyVaultKeyId", GoMethod: "ResetManagedServicesCmkKeyVaultKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSecurityGroupRulesRequired", GoMethod: "ResetNetworkSecurityGroupRulesRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicNetworkAccessEnabled", GoMethod: "ResetPublicNetworkAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sku", GoGetter: "Sku"},
			_jsii_.MemberProperty{JsiiProperty: "skuInput", GoGetter: "SkuInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountIdentity", GoGetter: "StorageAccountIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceUrl", GoGetter: "WorkspaceUrl"},
		},
		func() interface{} {
			j := jsiiProxy_DatabricksWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.databricksWorkspace.DatabricksWorkspaceConfig",
		reflect.TypeOf((*DatabricksWorkspaceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.databricksWorkspace.DatabricksWorkspaceCustomParameters",
		reflect.TypeOf((*DatabricksWorkspaceCustomParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.databricksWorkspace.DatabricksWorkspaceCustomParametersOutputReference",
		reflect.TypeOf((*DatabricksWorkspaceCustomParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "machineLearningWorkspaceId", GoGetter: "MachineLearningWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "machineLearningWorkspaceIdInput", GoGetter: "MachineLearningWorkspaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayName", GoGetter: "NatGatewayName"},
			_jsii_.MemberProperty{JsiiProperty: "natGatewayNameInput", GoGetter: "NatGatewayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "noPublicIp", GoGetter: "NoPublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "noPublicIpInput", GoGetter: "NoPublicIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetName", GoGetter: "PrivateSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetNameInput", GoGetter: "PrivateSubnetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetNetworkSecurityGroupAssociationId", GoGetter: "PrivateSubnetNetworkSecurityGroupAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "privateSubnetNetworkSecurityGroupAssociationIdInput", GoGetter: "PrivateSubnetNetworkSecurityGroupAssociationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpName", GoGetter: "PublicIpName"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpNameInput", GoGetter: "PublicIpNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetName", GoGetter: "PublicSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetNameInput", GoGetter: "PublicSubnetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetNetworkSecurityGroupAssociationId", GoGetter: "PublicSubnetNetworkSecurityGroupAssociationId"},
			_jsii_.MemberProperty{JsiiProperty: "publicSubnetNetworkSecurityGroupAssociationIdInput", GoGetter: "PublicSubnetNetworkSecurityGroupAssociationIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMachineLearningWorkspaceId", GoMethod: "ResetMachineLearningWorkspaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNatGatewayName", GoMethod: "ResetNatGatewayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoPublicIp", GoMethod: "ResetNoPublicIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateSubnetName", GoMethod: "ResetPrivateSubnetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateSubnetNetworkSecurityGroupAssociationId", GoMethod: "ResetPrivateSubnetNetworkSecurityGroupAssociationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicIpName", GoMethod: "ResetPublicIpName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicSubnetName", GoMethod: "ResetPublicSubnetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicSubnetNetworkSecurityGroupAssociationId", GoMethod: "ResetPublicSubnetNetworkSecurityGroupAssociationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountName", GoMethod: "ResetStorageAccountName"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageAccountSkuName", GoMethod: "ResetStorageAccountSkuName"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualNetworkId", GoMethod: "ResetVirtualNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVnetAddressPrefix", GoMethod: "ResetVnetAddressPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountName", GoGetter: "StorageAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountNameInput", GoGetter: "StorageAccountNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountSkuName", GoGetter: "StorageAccountSkuName"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountSkuNameInput", GoGetter: "StorageAccountSkuNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkIdInput", GoGetter: "VirtualNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vnetAddressPrefix", GoGetter: "VnetAddressPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "vnetAddressPrefixInput", GoGetter: "VnetAddressPrefixInput"},
		},
		func() interface{} {
			j := jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentity",
		reflect.TypeOf((*DatabricksWorkspaceStorageAccountIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityList",
		reflect.TypeOf((*DatabricksWorkspaceStorageAccountIdentityList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DatabricksWorkspaceStorageAccountIdentityList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.databricksWorkspace.DatabricksWorkspaceStorageAccountIdentityOutputReference",
		reflect.TypeOf((*DatabricksWorkspaceStorageAccountIdentityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "principalId", GoGetter: "PrincipalId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_DatabricksWorkspaceStorageAccountIdentityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.databricksWorkspace.DatabricksWorkspaceTimeouts",
		reflect.TypeOf((*DatabricksWorkspaceTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.databricksWorkspace.DatabricksWorkspaceTimeoutsOutputReference",
		reflect.TypeOf((*DatabricksWorkspaceTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_DatabricksWorkspaceTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
