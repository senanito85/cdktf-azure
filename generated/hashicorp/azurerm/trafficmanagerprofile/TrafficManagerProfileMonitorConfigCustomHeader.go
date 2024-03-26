package trafficmanagerprofile


type TrafficManagerProfileMonitorConfigCustomHeader struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_profile#name TrafficManagerProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_profile#value TrafficManagerProfile#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

