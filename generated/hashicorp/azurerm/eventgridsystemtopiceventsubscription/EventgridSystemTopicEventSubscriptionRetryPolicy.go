package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_system_topic_event_subscription#event_time_to_live EventgridSystemTopicEventSubscription#event_time_to_live}.
	EventTimeToLive *float64 `field:"required" json:"eventTimeToLive" yaml:"eventTimeToLive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_system_topic_event_subscription#max_delivery_attempts EventgridSystemTopicEventSubscription#max_delivery_attempts}.
	MaxDeliveryAttempts *float64 `field:"required" json:"maxDeliveryAttempts" yaml:"maxDeliveryAttempts"`
}

