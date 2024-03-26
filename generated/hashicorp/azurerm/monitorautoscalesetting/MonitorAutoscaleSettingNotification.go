package monitorautoscalesetting


type MonitorAutoscaleSettingNotification struct {
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_autoscale_setting#email MonitorAutoscaleSetting#email}
	Email *MonitorAutoscaleSettingNotificationEmail `field:"optional" json:"email" yaml:"email"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_autoscale_setting#webhook MonitorAutoscaleSetting#webhook}
	Webhook interface{} `field:"optional" json:"webhook" yaml:"webhook"`
}

