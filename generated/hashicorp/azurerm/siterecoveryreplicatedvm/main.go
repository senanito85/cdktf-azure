package siterecoveryreplicatedvm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		reflect.TypeOf((*SiteRecoveryReplicatedVm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "managedDisk", GoGetter: "ManagedDisk"},
			_jsii_.MemberProperty{JsiiProperty: "managedDiskInput", GoGetter: "ManagedDiskInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterface", GoGetter: "NetworkInterface"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceInput", GoGetter: "NetworkInterfaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putManagedDisk", GoMethod: "PutManagedDisk"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterface", GoMethod: "PutNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryReplicationPolicyId", GoGetter: "RecoveryReplicationPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryReplicationPolicyIdInput", GoGetter: "RecoveryReplicationPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryVaultName", GoGetter: "RecoveryVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryVaultNameInput", GoGetter: "RecoveryVaultNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedDisk", GoMethod: "ResetManagedDisk"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterface", GoMethod: "ResetNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetAvailabilitySetId", GoMethod: "ResetTargetAvailabilitySetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetNetworkId", GoMethod: "ResetTargetNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryFabricName", GoGetter: "SourceRecoveryFabricName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryFabricNameInput", GoGetter: "SourceRecoveryFabricNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryProtectionContainerName", GoGetter: "SourceRecoveryProtectionContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryProtectionContainerNameInput", GoGetter: "SourceRecoveryProtectionContainerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVmId", GoGetter: "SourceVmId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVmIdInput", GoGetter: "SourceVmIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetAvailabilitySetId", GoGetter: "TargetAvailabilitySetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetAvailabilitySetIdInput", GoGetter: "TargetAvailabilitySetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworkId", GoGetter: "TargetNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworkIdInput", GoGetter: "TargetNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryFabricId", GoGetter: "TargetRecoveryFabricId"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryFabricIdInput", GoGetter: "TargetRecoveryFabricIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryProtectionContainerId", GoGetter: "TargetRecoveryProtectionContainerId"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryProtectionContainerIdInput", GoGetter: "TargetRecoveryProtectionContainerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupId", GoGetter: "TargetResourceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupIdInput", GoGetter: "TargetResourceGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVm{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmConfig",
		reflect.TypeOf((*SiteRecoveryReplicatedVmConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDisk",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskId", GoGetter: "DiskId"},
			_jsii_.MemberProperty{JsiiProperty: "diskIdInput", GoGetter: "DiskIdInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDiskId", GoMethod: "ResetDiskId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStagingStorageAccountId", GoMethod: "ResetStagingStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDiskEncryptionSetId", GoMethod: "ResetTargetDiskEncryptionSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDiskType", GoMethod: "ResetTargetDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetReplicaDiskType", GoMethod: "ResetTargetReplicaDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetResourceGroupId", GoMethod: "ResetTargetResourceGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountId", GoGetter: "StagingStorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountIdInput", GoGetter: "StagingStorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryptionSetId", GoGetter: "TargetDiskEncryptionSetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryptionSetIdInput", GoGetter: "TargetDiskEncryptionSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskType", GoGetter: "TargetDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskTypeInput", GoGetter: "TargetDiskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetReplicaDiskType", GoGetter: "TargetReplicaDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "targetReplicaDiskTypeInput", GoGetter: "TargetReplicaDiskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupId", GoGetter: "TargetResourceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupIdInput", GoGetter: "TargetResourceGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterface",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterfaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "recoveryPublicIpAddressId", GoGetter: "RecoveryPublicIpAddressId"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryPublicIpAddressIdInput", GoGetter: "RecoveryPublicIpAddressIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoveryPublicIpAddressId", GoMethod: "ResetRecoveryPublicIpAddressId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceNetworkInterfaceId", GoMethod: "ResetSourceNetworkInterfaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetStaticIp", GoMethod: "ResetTargetStaticIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSubnetName", GoMethod: "ResetTargetSubnetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNetworkInterfaceId", GoGetter: "SourceNetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNetworkInterfaceIdInput", GoGetter: "SourceNetworkInterfaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetStaticIp", GoGetter: "TargetStaticIp"},
			_jsii_.MemberProperty{JsiiProperty: "targetStaticIpInput", GoGetter: "TargetStaticIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetSubnetName", GoGetter: "TargetSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "targetSubnetNameInput", GoGetter: "TargetSubnetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmTimeouts",
		reflect.TypeOf((*SiteRecoveryReplicatedVmTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmTimeoutsOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_SiteRecoveryReplicatedVmTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
