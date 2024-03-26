package apimanagement


type ApiManagementAdditionalLocationVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management#subnet_id ApiManagement#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

