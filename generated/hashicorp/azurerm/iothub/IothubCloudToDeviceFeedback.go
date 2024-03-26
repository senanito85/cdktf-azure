package iothub


type IothubCloudToDeviceFeedback struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#lock_duration Iothub#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#max_delivery_count Iothub#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"optional" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#time_to_live Iothub#time_to_live}.
	TimeToLive *string `field:"optional" json:"timeToLive" yaml:"timeToLive"`
}

