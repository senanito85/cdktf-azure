package networkwatcherflowlog


type NetworkWatcherFlowLogTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_watcher_flow_log#create NetworkWatcherFlowLog#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_watcher_flow_log#delete NetworkWatcherFlowLog#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_watcher_flow_log#read NetworkWatcherFlowLog#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_watcher_flow_log#update NetworkWatcherFlowLog#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

