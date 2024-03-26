package monitormetricalert


type MonitorMetricAlertAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_metric_alert#action_group_id MonitorMetricAlert#action_group_id}.
	ActionGroupId *string `field:"required" json:"actionGroupId" yaml:"actionGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_metric_alert#webhook_properties MonitorMetricAlert#webhook_properties}.
	WebhookProperties *map[string]*string `field:"optional" json:"webhookProperties" yaml:"webhookProperties"`
}

