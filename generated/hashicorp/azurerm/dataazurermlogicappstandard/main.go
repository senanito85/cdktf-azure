package dataazurermlogicappstandard

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandard",
		reflect.TypeOf((*DataAzurermLogicAppStandard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appServicePlanId", GoGetter: "AppServicePlanId"},
			_jsii_.MemberProperty{JsiiProperty: "appSettings", GoGetter: "AppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "bundleVersion", GoGetter: "BundleVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientAffinityEnabled", GoGetter: "ClientAffinityEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificateMode", GoGetter: "ClientCertificateMode"},
			_jsii_.MemberProperty{JsiiProperty: "connectionString", GoGetter: "ConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customDomainVerificationId", GoGetter: "CustomDomainVerificationId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultHostname", GoGetter: "DefaultHostname"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpsOnly", GoGetter: "HttpsOnly"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identity", GoGetter: "Identity"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outboundIpAddresses", GoGetter: "OutboundIpAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "possibleOutboundIpAddresses", GoGetter: "PossibleOutboundIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSiteConfig", GoMethod: "PutSiteConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSiteConfig", GoMethod: "ResetSiteConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "siteConfig", GoGetter: "SiteConfig"},
			_jsii_.MemberProperty{JsiiProperty: "siteConfigInput", GoGetter: "SiteConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "siteCredential", GoGetter: "SiteCredential"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountAccessKey", GoGetter: "StorageAccountAccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountName", GoGetter: "StorageAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "storageAccountShareName", GoGetter: "StorageAccountShareName"},
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
			_jsii_.MemberProperty{JsiiProperty: "useExtensionBundle", GoGetter: "UseExtensionBundle"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandard{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardConfig",
		reflect.TypeOf((*DataAzurermLogicAppStandardConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardConnectionString",
		reflect.TypeOf((*DataAzurermLogicAppStandardConnectionString)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardConnectionStringList",
		reflect.TypeOf((*DataAzurermLogicAppStandardConnectionStringList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardConnectionStringList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardConnectionStringOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardConnectionStringOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardConnectionStringOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardIdentity",
		reflect.TypeOf((*DataAzurermLogicAppStandardIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardIdentityList",
		reflect.TypeOf((*DataAzurermLogicAppStandardIdentityList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardIdentityList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardIdentityOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardIdentityOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardIdentityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfig",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigCors",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigCors)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigCorsOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigCorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedOrigins", GoGetter: "AllowedOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOriginsInput", GoGetter: "AllowedOriginsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSupportCredentials", GoMethod: "ResetSupportCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "supportCredentials", GoGetter: "SupportCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "supportCredentialsInput", GoGetter: "SupportCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigCorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestriction",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestriction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestrictionHeaders",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestrictionHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersList",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetXAzureFdid", GoMethod: "ResetXAzureFdid"},
			_jsii_.MemberMethod{JsiiMethod: "resetXFdHealthProbe", GoMethod: "ResetXFdHealthProbe"},
			_jsii_.MemberMethod{JsiiMethod: "resetXForwardedFor", GoMethod: "ResetXForwardedFor"},
			_jsii_.MemberMethod{JsiiMethod: "resetXForwardedHost", GoMethod: "ResetXForwardedHost"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "xAzureFdid", GoGetter: "XAzureFdid"},
			_jsii_.MemberProperty{JsiiProperty: "xAzureFdidInput", GoGetter: "XAzureFdidInput"},
			_jsii_.MemberProperty{JsiiProperty: "xFdHealthProbe", GoGetter: "XFdHealthProbe"},
			_jsii_.MemberProperty{JsiiProperty: "xFdHealthProbeInput", GoGetter: "XFdHealthProbeInput"},
			_jsii_.MemberProperty{JsiiProperty: "xForwardedFor", GoGetter: "XForwardedFor"},
			_jsii_.MemberProperty{JsiiProperty: "xForwardedForInput", GoGetter: "XForwardedForInput"},
			_jsii_.MemberProperty{JsiiProperty: "xForwardedHost", GoGetter: "XForwardedHost"},
			_jsii_.MemberProperty{JsiiProperty: "xForwardedHostInput", GoGetter: "XForwardedHostInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigIpRestrictionHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestrictionList",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestrictionList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigIpRestrictionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigIpRestrictionOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigIpRestrictionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressInput", GoGetter: "IpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaders", GoMethod: "PutHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAction", GoMethod: "ResetAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpAddress", GoMethod: "ResetIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceTag", GoMethod: "ResetServiceTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualNetworkSubnetId", GoMethod: "ResetVirtualNetworkSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceTag", GoGetter: "ServiceTag"},
			_jsii_.MemberProperty{JsiiProperty: "serviceTagInput", GoGetter: "ServiceTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkSubnetId", GoGetter: "VirtualNetworkSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkSubnetIdInput", GoGetter: "VirtualNetworkSubnetIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigIpRestrictionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteConfigOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysOn", GoGetter: "AlwaysOn"},
			_jsii_.MemberProperty{JsiiProperty: "alwaysOnInput", GoGetter: "AlwaysOnInput"},
			_jsii_.MemberProperty{JsiiProperty: "appScaleLimit", GoGetter: "AppScaleLimit"},
			_jsii_.MemberProperty{JsiiProperty: "appScaleLimitInput", GoGetter: "AppScaleLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cors", GoGetter: "Cors"},
			_jsii_.MemberProperty{JsiiProperty: "corsInput", GoGetter: "CorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dotnetFrameworkVersion", GoGetter: "DotnetFrameworkVersion"},
			_jsii_.MemberProperty{JsiiProperty: "dotnetFrameworkVersionInput", GoGetter: "DotnetFrameworkVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticInstanceMinimum", GoGetter: "ElasticInstanceMinimum"},
			_jsii_.MemberProperty{JsiiProperty: "elasticInstanceMinimumInput", GoGetter: "ElasticInstanceMinimumInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "ftpsState", GoGetter: "FtpsState"},
			_jsii_.MemberProperty{JsiiProperty: "ftpsStateInput", GoGetter: "FtpsStateInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckPath", GoGetter: "HealthCheckPath"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckPathInput", GoGetter: "HealthCheckPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "http2Enabled", GoGetter: "Http2Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "http2EnabledInput", GoGetter: "Http2EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipRestriction", GoGetter: "IpRestriction"},
			_jsii_.MemberProperty{JsiiProperty: "ipRestrictionInput", GoGetter: "IpRestrictionInput"},
			_jsii_.MemberProperty{JsiiProperty: "linuxFxVersion", GoGetter: "LinuxFxVersion"},
			_jsii_.MemberProperty{JsiiProperty: "linuxFxVersionInput", GoGetter: "LinuxFxVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersion", GoGetter: "MinTlsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minTlsVersionInput", GoGetter: "MinTlsVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "preWarmedInstanceCount", GoGetter: "PreWarmedInstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "preWarmedInstanceCountInput", GoGetter: "PreWarmedInstanceCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCors", GoMethod: "PutCors"},
			_jsii_.MemberMethod{JsiiMethod: "putIpRestriction", GoMethod: "PutIpRestriction"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlwaysOn", GoMethod: "ResetAlwaysOn"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppScaleLimit", GoMethod: "ResetAppScaleLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetCors", GoMethod: "ResetCors"},
			_jsii_.MemberMethod{JsiiMethod: "resetDotnetFrameworkVersion", GoMethod: "ResetDotnetFrameworkVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticInstanceMinimum", GoMethod: "ResetElasticInstanceMinimum"},
			_jsii_.MemberMethod{JsiiMethod: "resetFtpsState", GoMethod: "ResetFtpsState"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckPath", GoMethod: "ResetHealthCheckPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2Enabled", GoMethod: "ResetHttp2Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpRestriction", GoMethod: "ResetIpRestriction"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinuxFxVersion", GoMethod: "ResetLinuxFxVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTlsVersion", GoMethod: "ResetMinTlsVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreWarmedInstanceCount", GoMethod: "ResetPreWarmedInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetRuntimeScaleMonitoringEnabled", GoMethod: "ResetRuntimeScaleMonitoringEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetUse32BitWorkerProcess", GoMethod: "ResetUse32BitWorkerProcess"},
			_jsii_.MemberMethod{JsiiMethod: "resetVnetRouteAllEnabled", GoMethod: "ResetVnetRouteAllEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebsocketsEnabled", GoMethod: "ResetWebsocketsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeScaleMonitoringEnabled", GoGetter: "RuntimeScaleMonitoringEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeScaleMonitoringEnabledInput", GoGetter: "RuntimeScaleMonitoringEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "use32BitWorkerProcess", GoGetter: "Use32BitWorkerProcess"},
			_jsii_.MemberProperty{JsiiProperty: "use32BitWorkerProcessInput", GoGetter: "Use32BitWorkerProcessInput"},
			_jsii_.MemberProperty{JsiiProperty: "vnetRouteAllEnabled", GoGetter: "VnetRouteAllEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "vnetRouteAllEnabledInput", GoGetter: "VnetRouteAllEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "websocketsEnabled", GoGetter: "WebsocketsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "websocketsEnabledInput", GoGetter: "WebsocketsEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardSiteConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteCredential",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteCredential)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteCredentialList",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteCredentialList)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardSiteCredentialList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardSiteCredentialOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardSiteCredentialOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			j := jsiiProxy_DataAzurermLogicAppStandardSiteCredentialOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardTimeouts",
		reflect.TypeOf((*DataAzurermLogicAppStandardTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"azurerm.dataAzurermLogicAppStandard.DataAzurermLogicAppStandardTimeoutsOutputReference",
		reflect.TypeOf((*DataAzurermLogicAppStandardTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_DataAzurermLogicAppStandardTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
