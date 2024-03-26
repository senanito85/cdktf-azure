package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyManagedRulesManagedRuleSetRuleGroupOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_application_firewall_policy#rule_group_name WebApplicationFirewallPolicy#rule_group_name}.
	RuleGroupName *string `field:"required" json:"ruleGroupName" yaml:"ruleGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_application_firewall_policy#disabled_rules WebApplicationFirewallPolicy#disabled_rules}.
	DisabledRules *[]*string `field:"optional" json:"disabledRules" yaml:"disabledRules"`
}

