package networkconnectionmonitor


type NetworkConnectionMonitorSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#port NetworkConnectionMonitor#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_connection_monitor#virtual_machine_id NetworkConnectionMonitor#virtual_machine_id}.
	VirtualMachineId *string `field:"optional" json:"virtualMachineId" yaml:"virtualMachineId"`
}

