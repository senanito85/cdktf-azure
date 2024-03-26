package apimanagementapidiagnostic


type ApiManagementApiDiagnosticBackendRequestDataMaskingHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_diagnostic#mode ApiManagementApiDiagnostic#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_diagnostic#value ApiManagementApiDiagnostic#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

