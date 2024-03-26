package eventhubnamespace


type EventhubNamespaceNetworkRulesets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace#default_action EventhubNamespace#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace#ip_rule EventhubNamespace#ip_rule}.
	IpRule interface{} `field:"optional" json:"ipRule" yaml:"ipRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace#trusted_service_access_enabled EventhubNamespace#trusted_service_access_enabled}.
	TrustedServiceAccessEnabled interface{} `field:"optional" json:"trustedServiceAccessEnabled" yaml:"trustedServiceAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace#virtual_network_rule EventhubNamespace#virtual_network_rule}.
	VirtualNetworkRule interface{} `field:"optional" json:"virtualNetworkRule" yaml:"virtualNetworkRule"`
}

