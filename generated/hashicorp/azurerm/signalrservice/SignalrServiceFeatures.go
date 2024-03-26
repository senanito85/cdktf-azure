package signalrservice


type SignalrServiceFeatures struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service#flag SignalrService#flag}.
	Flag *string `field:"required" json:"flag" yaml:"flag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service#value SignalrService#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

