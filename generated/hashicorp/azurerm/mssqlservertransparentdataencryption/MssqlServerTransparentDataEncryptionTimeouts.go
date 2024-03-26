package mssqlservertransparentdataencryption


type MssqlServerTransparentDataEncryptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_transparent_data_encryption#create MssqlServerTransparentDataEncryption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_transparent_data_encryption#delete MssqlServerTransparentDataEncryption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_transparent_data_encryption#read MssqlServerTransparentDataEncryption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server_transparent_data_encryption#update MssqlServerTransparentDataEncryption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

