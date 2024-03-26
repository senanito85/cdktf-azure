package mariadbserver


type MariadbServerStorageProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mariadb_server#auto_grow MariadbServer#auto_grow}.
	AutoGrow *string `field:"optional" json:"autoGrow" yaml:"autoGrow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mariadb_server#backup_retention_days MariadbServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mariadb_server#geo_redundant_backup MariadbServer#geo_redundant_backup}.
	GeoRedundantBackup *string `field:"optional" json:"geoRedundantBackup" yaml:"geoRedundantBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mariadb_server#storage_mb MariadbServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
}

