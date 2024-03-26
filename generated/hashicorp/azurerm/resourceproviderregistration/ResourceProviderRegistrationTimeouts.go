package resourceproviderregistration


type ResourceProviderRegistrationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/resource_provider_registration#create ResourceProviderRegistration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/resource_provider_registration#delete ResourceProviderRegistration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/resource_provider_registration#read ResourceProviderRegistration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/resource_provider_registration#update ResourceProviderRegistration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

