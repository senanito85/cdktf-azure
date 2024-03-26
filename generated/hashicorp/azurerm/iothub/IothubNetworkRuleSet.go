package iothub


type IothubNetworkRuleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#apply_to_builtin_eventhub_endpoint Iothub#apply_to_builtin_eventhub_endpoint}.
	ApplyToBuiltinEventhubEndpoint interface{} `field:"optional" json:"applyToBuiltinEventhubEndpoint" yaml:"applyToBuiltinEventhubEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#default_action Iothub#default_action}.
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// ip_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#ip_rule Iothub#ip_rule}
	IpRule interface{} `field:"optional" json:"ipRule" yaml:"ipRule"`
}

