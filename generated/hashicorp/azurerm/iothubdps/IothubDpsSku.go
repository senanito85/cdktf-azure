package iothubdps


type IothubDpsSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_dps#capacity IothubDps#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_dps#name IothubDps#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

