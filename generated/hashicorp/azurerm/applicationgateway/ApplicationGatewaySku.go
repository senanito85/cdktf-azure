package applicationgateway


type ApplicationGatewaySku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#tier ApplicationGateway#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#capacity ApplicationGateway#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

