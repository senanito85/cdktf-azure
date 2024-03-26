package monitorsmartdetectoralertrule


type MonitorSmartDetectorAlertRuleActionGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_smart_detector_alert_rule#ids MonitorSmartDetectorAlertRule#ids}.
	Ids *[]*string `field:"required" json:"ids" yaml:"ids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_smart_detector_alert_rule#email_subject MonitorSmartDetectorAlertRule#email_subject}.
	EmailSubject *string `field:"optional" json:"emailSubject" yaml:"emailSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_smart_detector_alert_rule#webhook_payload MonitorSmartDetectorAlertRule#webhook_payload}.
	WebhookPayload *string `field:"optional" json:"webhookPayload" yaml:"webhookPayload"`
}

