package monitordiagnosticsetting


type MonitorDiagnosticSettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#create MonitorDiagnosticSetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#delete MonitorDiagnosticSetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#read MonitorDiagnosticSetting#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_diagnostic_setting#update MonitorDiagnosticSetting#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

