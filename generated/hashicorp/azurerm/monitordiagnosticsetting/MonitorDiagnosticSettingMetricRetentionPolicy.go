package monitordiagnosticsetting


type MonitorDiagnosticSettingMetricRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#enabled MonitorDiagnosticSetting#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#days MonitorDiagnosticSetting#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

