package datafactorylinkedserviceazuresqldatabase


type DataFactoryLinkedServiceAzureSqlDatabaseKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_azure_sql_database#linked_service_name DataFactoryLinkedServiceAzureSqlDatabase#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_azure_sql_database#secret_name DataFactoryLinkedServiceAzureSqlDatabase#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

