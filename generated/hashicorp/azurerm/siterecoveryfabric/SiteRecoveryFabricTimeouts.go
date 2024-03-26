package siterecoveryfabric


type SiteRecoveryFabricTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_fabric#create SiteRecoveryFabric#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_fabric#delete SiteRecoveryFabric#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_fabric#read SiteRecoveryFabric#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_fabric#update SiteRecoveryFabric#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

