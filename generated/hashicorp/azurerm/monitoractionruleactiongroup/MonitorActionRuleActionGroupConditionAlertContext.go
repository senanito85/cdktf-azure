package monitoractionruleactiongroup


type MonitorActionRuleActionGroupConditionAlertContext struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_action_rule_action_group#operator MonitorActionRuleActionGroup#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_action_rule_action_group#values MonitorActionRuleActionGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

