package keyvaultmanagedstorageaccount


type KeyVaultManagedStorageAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_managed_storage_account#create KeyVaultManagedStorageAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_managed_storage_account#delete KeyVaultManagedStorageAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_managed_storage_account#read KeyVaultManagedStorageAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_managed_storage_account#update KeyVaultManagedStorageAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

