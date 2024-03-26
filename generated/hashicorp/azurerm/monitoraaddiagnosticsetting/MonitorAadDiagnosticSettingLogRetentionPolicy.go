package monitoraaddiagnosticsetting


type MonitorAadDiagnosticSettingLogRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_aad_diagnostic_setting#days MonitorAadDiagnosticSetting#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_aad_diagnostic_setting#enabled MonitorAadDiagnosticSetting#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

