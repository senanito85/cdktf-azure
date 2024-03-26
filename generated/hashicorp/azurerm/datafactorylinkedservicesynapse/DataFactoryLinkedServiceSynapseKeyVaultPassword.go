package datafactorylinkedservicesynapse


type DataFactoryLinkedServiceSynapseKeyVaultPassword struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_synapse#linked_service_name DataFactoryLinkedServiceSynapse#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_synapse#secret_name DataFactoryLinkedServiceSynapse#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
}

