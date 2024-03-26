package datafactorylinkedserviceazureblobstorage


type DataFactoryLinkedServiceAzureBlobStorageKeyVaultSasToken struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_azure_blob_storage#linked_service_name DataFactoryLinkedServiceAzureBlobStorage#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_azure_blob_storage#secret_name DataFactoryLinkedServiceAzureBlobStorage#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

