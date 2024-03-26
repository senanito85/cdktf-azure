package mssqldatabase


type MssqlDatabaseThreatDetectionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#disabled_alerts MssqlDatabase#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#email_account_admins MssqlDatabase#email_account_admins}.
	EmailAccountAdmins *string `field:"optional" json:"emailAccountAdmins" yaml:"emailAccountAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#email_addresses MssqlDatabase#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#state MssqlDatabase#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#storage_account_access_key MssqlDatabase#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#storage_endpoint MssqlDatabase#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#use_server_default MssqlDatabase#use_server_default}.
	UseServerDefault *string `field:"optional" json:"useServerDefault" yaml:"useServerDefault"`
}

