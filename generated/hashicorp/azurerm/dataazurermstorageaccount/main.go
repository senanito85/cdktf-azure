package dataazurermstorageaccount

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccount",
		reflect.TypeOf((*DataAzurermStorageAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTier", GoGetter: "AccessTier"},
			_jsii_.MemberProperty{JsiiProperty: "accountKind", GoGetter: "AccountKind"},
			_jsii_.MemberProperty{JsiiProperty: "accountReplicationType", GoGetter: "AccountReplicationType"},
			_jsii_.MemberProperty{JsiiProperty: "accountTier", GoGetter: "AccountTier"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowBlobPublicAccess", GoGetter: "AllowBlobPublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customDomain", GoGetter: "CustomDomain"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableHttpsTrafficOnly", GoGetter: "EnableHttpsTrafficOnly"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureEncryptionEnabled", GoGetter: "InfrastructureEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isHnsEnabled", GoGetter: "IsHnsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersion", GoGetter: "MinTlsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersionInput", GoGetter: "MinTlsVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryAccessKey", GoGetter: "PrimaryAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobConnectionString", GoGetter: "PrimaryBlobConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobEndpoint", GoGetter: "PrimaryBlobEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryBlobHost", GoGetter: "PrimaryBlobHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryConnectionString", GoGetter: "PrimaryConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsEndpoint", GoGetter: "PrimaryDfsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDfsHost", GoGetter: "PrimaryDfsHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileEndpoint", GoGetter: "PrimaryFileEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryFileHost", GoGetter: "PrimaryFileHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryLocation", GoGetter: "PrimaryLocation"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueEndpoint", GoGetter: "PrimaryQueueEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryQueueHost", GoGetter: "PrimaryQueueHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableEndpoint", GoGetter: "PrimaryTableEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryTableHost", GoGetter: "PrimaryTableHost"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebEndpoint", GoGetter: "PrimaryWebEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "primaryWebHost", GoGetter: "PrimaryWebHost"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "queueEncryptionKeyType", GoGetter: "QueueEncryptionKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTlsVersion", GoMethod: "ResetMinTlsVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryAccessKey", GoGetter: "SecondaryAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobConnectionString", GoGetter: "SecondaryBlobConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobEndpoint", GoGetter: "SecondaryBlobEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryBlobHost", GoGetter: "SecondaryBlobHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryConnectionString", GoGetter: "SecondaryConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsEndpoint", GoGetter: "SecondaryDfsEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryDfsHost", GoGetter: "SecondaryDfsHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileEndpoint", GoGetter: "SecondaryFileEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryFileHost", GoGetter: "SecondaryFileHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryLocation", GoGetter: "SecondaryLocation"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueEndpoint", GoGetter: "SecondaryQueueEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryQueueHost", GoGetter: "SecondaryQueueHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableEndpoint", GoGetter: "SecondaryTableEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryTableHost", GoGetter: "SecondaryTableHost"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebEndpoint", GoGetter: "SecondaryWebEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWebHost", GoGetter: "SecondaryWebHost"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tableEncryptionKeyType", GoGetter: "TableEncryptionKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
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
			j := jsiiProxy_DataAzurermStorageAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountConfig",
		reflect.TypeOf((*DataAzurermStorageAccountConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomain",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomain)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomainList",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomainList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermStorageAccountCustomDomainList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountCustomDomainOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountCustomDomainOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountCustomDomainOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountTimeouts",
		reflect.TypeOf((*DataAzurermStorageAccountTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermStorageAccount.DataAzurermStorageAccountTimeoutsOutputReference",
		reflect.TypeOf((*DataAzurermStorageAccountTimeoutsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermStorageAccountTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
