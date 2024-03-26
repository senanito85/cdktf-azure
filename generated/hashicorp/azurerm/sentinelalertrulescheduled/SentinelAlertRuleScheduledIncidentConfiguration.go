package sentinelalertrulescheduled


type SentinelAlertRuleScheduledIncidentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#create_incident SentinelAlertRuleScheduled#create_incident}.
	CreateIncident interface{} `field:"required" json:"createIncident" yaml:"createIncident"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#grouping SentinelAlertRuleScheduled#grouping}
	Grouping *SentinelAlertRuleScheduledIncidentConfigurationGrouping `field:"required" json:"grouping" yaml:"grouping"`
}

