package apimanagement


type ApiManagementAdditionalLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#location ApiManagement#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#capacity ApiManagement#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#public_ip_address_id ApiManagement#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#virtual_network_configuration ApiManagement#virtual_network_configuration}
	VirtualNetworkConfiguration *ApiManagementAdditionalLocationVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#zones ApiManagement#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

