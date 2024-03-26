package mysqlserver


type MysqlServerStorageProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mysql_server#auto_grow MysqlServer#auto_grow}.
	AutoGrow *string `field:"optional" json:"autoGrow" yaml:"autoGrow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mysql_server#backup_retention_days MysqlServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mysql_server#geo_redundant_backup MysqlServer#geo_redundant_backup}.
	GeoRedundantBackup *string `field:"optional" json:"geoRedundantBackup" yaml:"geoRedundantBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mysql_server#storage_mb MysqlServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
}

