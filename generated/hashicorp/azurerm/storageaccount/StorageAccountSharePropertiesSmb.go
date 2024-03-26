package storageaccount


type StorageAccountSharePropertiesSmb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account#authentication_types StorageAccount#authentication_types}.
	AuthenticationTypes *[]*string `field:"optional" json:"authenticationTypes" yaml:"authenticationTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account#channel_encryption_type StorageAccount#channel_encryption_type}.
	ChannelEncryptionType *[]*string `field:"optional" json:"channelEncryptionType" yaml:"channelEncryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account#kerberos_ticket_encryption_type StorageAccount#kerberos_ticket_encryption_type}.
	KerberosTicketEncryptionType *[]*string `field:"optional" json:"kerberosTicketEncryptionType" yaml:"kerberosTicketEncryptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account#versions StorageAccount#versions}.
	Versions *[]*string `field:"optional" json:"versions" yaml:"versions"`
}

