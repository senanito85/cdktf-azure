package mssqldatabase


type MssqlDatabaseShortTermRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_database#retention_days MssqlDatabase#retention_days}.
	RetentionDays *float64 `field:"required" json:"retentionDays" yaml:"retentionDays"`
}

