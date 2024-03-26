package frontdoorfirewallpolicy


type FrontdoorFirewallPolicyManagedRuleOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/frontdoor_firewall_policy#rule_group_name FrontdoorFirewallPolicy#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/frontdoor_firewall_policy#exclusion FrontdoorFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/frontdoor_firewall_policy#rule FrontdoorFirewallPolicy#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

