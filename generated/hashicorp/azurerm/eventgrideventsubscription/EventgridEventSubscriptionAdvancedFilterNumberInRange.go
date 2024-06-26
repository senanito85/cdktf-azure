package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterNumberInRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values interface{} `field:"required" json:"values" yaml:"values"`
}

