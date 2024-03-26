package batchpool


type BatchPoolContainerConfigurationContainerRegistries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#password BatchPool#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#registry_server BatchPool#registry_server}.
	RegistryServer *string `field:"optional" json:"registryServer" yaml:"registryServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

