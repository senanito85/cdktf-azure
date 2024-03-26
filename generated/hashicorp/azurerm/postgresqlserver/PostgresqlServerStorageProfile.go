package postgresqlserver


type PostgresqlServerStorageProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/postgresql_server#auto_grow PostgresqlServer#auto_grow}.
	AutoGrow *string `field:"optional" json:"autoGrow" yaml:"autoGrow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/postgresql_server#backup_retention_days PostgresqlServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/postgresql_server#geo_redundant_backup PostgresqlServer#geo_redundant_backup}.
	GeoRedundantBackup *string `field:"optional" json:"geoRedundantBackup" yaml:"geoRedundantBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/postgresql_server#storage_mb PostgresqlServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
}

