package datafactorylinkedserviceodata


type DataFactoryLinkedServiceOdataBasicAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_odata#password DataFactoryLinkedServiceOdata#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_linked_service_odata#username DataFactoryLinkedServiceOdata#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

