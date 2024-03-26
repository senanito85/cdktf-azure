package siterecoverynetworkmapping


type SiteRecoveryNetworkMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_network_mapping#create SiteRecoveryNetworkMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_network_mapping#delete SiteRecoveryNetworkMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_network_mapping#read SiteRecoveryNetworkMapping#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/site_recovery_network_mapping#update SiteRecoveryNetworkMapping#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

