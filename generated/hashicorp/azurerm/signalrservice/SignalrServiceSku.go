package signalrservice


type SignalrServiceSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service#capacity SignalrService#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service#name SignalrService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

