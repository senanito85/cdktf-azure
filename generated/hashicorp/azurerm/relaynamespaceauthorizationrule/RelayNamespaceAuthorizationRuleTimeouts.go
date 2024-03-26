package relaynamespaceauthorizationrule


type RelayNamespaceAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/relay_namespace_authorization_rule#create RelayNamespaceAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/relay_namespace_authorization_rule#delete RelayNamespaceAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/relay_namespace_authorization_rule#read RelayNamespaceAuthorizationRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/relay_namespace_authorization_rule#update RelayNamespaceAuthorizationRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

