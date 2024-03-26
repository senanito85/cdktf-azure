package apimanagementapioperation


type ApiManagementApiOperationRequestRepresentationExample struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation#name ApiManagementApiOperation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation#description ApiManagementApiOperation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation#external_value ApiManagementApiOperation#external_value}.
	ExternalValue *string `field:"optional" json:"externalValue" yaml:"externalValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation#summary ApiManagementApiOperation#summary}.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation#value ApiManagementApiOperation#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

