package apimanagementapitag


type ApiManagementApiTagTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_tag#create ApiManagementApiTag#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_tag#delete ApiManagementApiTag#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_tag#read ApiManagementApiTag#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

