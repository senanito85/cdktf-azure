package datafactoryintegrationruntimemanaged


type DataFactoryIntegrationRuntimeManagedCatalogInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_managed#server_endpoint DataFactoryIntegrationRuntimeManaged#server_endpoint}.
	ServerEndpoint *string `field:"required" json:"serverEndpoint" yaml:"serverEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_managed#administrator_login DataFactoryIntegrationRuntimeManaged#administrator_login}.
	AdministratorLogin *string `field:"optional" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_managed#administrator_password DataFactoryIntegrationRuntimeManaged#administrator_password}.
	AdministratorPassword *string `field:"optional" json:"administratorPassword" yaml:"administratorPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_managed#pricing_tier DataFactoryIntegrationRuntimeManaged#pricing_tier}.
	PricingTier *string `field:"optional" json:"pricingTier" yaml:"pricingTier"`
}

