package monitorautoscalesetting


type MonitorAutoscaleSettingProfileFixedDate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_autoscale_setting#end MonitorAutoscaleSetting#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_autoscale_setting#start MonitorAutoscaleSetting#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_autoscale_setting#timezone MonitorAutoscaleSetting#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

