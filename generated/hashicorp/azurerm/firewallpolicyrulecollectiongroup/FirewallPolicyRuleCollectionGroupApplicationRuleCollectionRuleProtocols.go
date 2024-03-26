package firewallpolicyrulecollectiongroup


type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRuleProtocols struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy_rule_collection_group#port FirewallPolicyRuleCollectionGroup#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy_rule_collection_group#type FirewallPolicyRuleCollectionGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

