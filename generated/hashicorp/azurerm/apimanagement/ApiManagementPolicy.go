package apimanagement


type ApiManagementPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#xml_content ApiManagement#xml_content}.
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#xml_link ApiManagement#xml_link}.
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

