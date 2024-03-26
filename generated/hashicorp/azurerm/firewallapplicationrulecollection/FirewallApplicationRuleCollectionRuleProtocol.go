package firewallapplicationrulecollection


type FirewallApplicationRuleCollectionRuleProtocol struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_application_rule_collection#port FirewallApplicationRuleCollection#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_application_rule_collection#type FirewallApplicationRuleCollection#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

