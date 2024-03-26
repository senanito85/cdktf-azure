package iothub


type IothubFallbackRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#condition Iothub#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#enabled Iothub#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#endpoint_names Iothub#endpoint_names}.
	EndpointNames *[]*string `field:"optional" json:"endpointNames" yaml:"endpointNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#source Iothub#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

