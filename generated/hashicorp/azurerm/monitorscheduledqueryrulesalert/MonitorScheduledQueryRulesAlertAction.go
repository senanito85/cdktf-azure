package monitorscheduledqueryrulesalert


type MonitorScheduledQueryRulesAlertAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_scheduled_query_rules_alert#action_group MonitorScheduledQueryRulesAlert#action_group}.
	ActionGroup *[]*string `field:"required" json:"actionGroup" yaml:"actionGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_scheduled_query_rules_alert#custom_webhook_payload MonitorScheduledQueryRulesAlert#custom_webhook_payload}.
	CustomWebhookPayload *string `field:"optional" json:"customWebhookPayload" yaml:"customWebhookPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_scheduled_query_rules_alert#email_subject MonitorScheduledQueryRulesAlert#email_subject}.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
}

