package manageddisk


type ManagedDiskEncryptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/managed_disk#enabled ManagedDisk#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/managed_disk#disk_encryption_key ManagedDisk#disk_encryption_key}
	DiskEncryptionKey *ManagedDiskEncryptionSettingsDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/managed_disk#key_encryption_key ManagedDisk#key_encryption_key}
	KeyEncryptionKey *ManagedDiskEncryptionSettingsKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

