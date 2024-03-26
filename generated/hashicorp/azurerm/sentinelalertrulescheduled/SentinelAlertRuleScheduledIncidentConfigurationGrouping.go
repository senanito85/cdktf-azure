package sentinelalertrulescheduled


type SentinelAlertRuleScheduledIncidentConfigurationGrouping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#entity_matching_method SentinelAlertRuleScheduled#entity_matching_method}.
	EntityMatchingMethod *string `field:"optional" json:"entityMatchingMethod" yaml:"entityMatchingMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#group_by SentinelAlertRuleScheduled#group_by}.
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#lookback_duration SentinelAlertRuleScheduled#lookback_duration}.
	LookbackDuration *string `field:"optional" json:"lookbackDuration" yaml:"lookbackDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#reopen_closed_incidents SentinelAlertRuleScheduled#reopen_closed_incidents}.
	ReopenClosedIncidents interface{} `field:"optional" json:"reopenClosedIncidents" yaml:"reopenClosedIncidents"`
}

