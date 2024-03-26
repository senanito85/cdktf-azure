package firewallpolicy


type FirewallPolicyIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#type FirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#user_assigned_identity_ids FirewallPolicy#user_assigned_identity_ids}.
	UserAssignedIdentityIds *[]*string `field:"optional" json:"userAssignedIdentityIds" yaml:"userAssignedIdentityIds"`
}

