package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyCustomRulesMatchConditionsMatchVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_application_firewall_policy#variable_name WebApplicationFirewallPolicy#variable_name}.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_application_firewall_policy#selector WebApplicationFirewallPolicy#selector}.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

