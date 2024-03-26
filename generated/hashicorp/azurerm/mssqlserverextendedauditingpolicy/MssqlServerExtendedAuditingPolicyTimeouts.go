package mssqlserverextendedauditingpolicy


type MssqlServerExtendedAuditingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_extended_auditing_policy#create MssqlServerExtendedAuditingPolicyA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_extended_auditing_policy#delete MssqlServerExtendedAuditingPolicyA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_extended_auditing_policy#read MssqlServerExtendedAuditingPolicyA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_extended_auditing_policy#update MssqlServerExtendedAuditingPolicyA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

