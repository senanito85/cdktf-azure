package networkconnectionmonitor


type NetworkConnectionMonitorTestGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#destination_endpoints NetworkConnectionMonitor#destination_endpoints}.
	DestinationEndpoints *[]*string `field:"required" json:"destinationEndpoints" yaml:"destinationEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#name NetworkConnectionMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#source_endpoints NetworkConnectionMonitor#source_endpoints}.
	SourceEndpoints *[]*string `field:"required" json:"sourceEndpoints" yaml:"sourceEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#test_configuration_names NetworkConnectionMonitor#test_configuration_names}.
	TestConfigurationNames *[]*string `field:"required" json:"testConfigurationNames" yaml:"testConfigurationNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#enabled NetworkConnectionMonitor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

