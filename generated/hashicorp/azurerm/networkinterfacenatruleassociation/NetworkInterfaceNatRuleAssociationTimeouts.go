package networkinterfacenatruleassociation


type NetworkInterfaceNatRuleAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_nat_rule_association#create NetworkInterfaceNatRuleAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_nat_rule_association#delete NetworkInterfaceNatRuleAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_nat_rule_association#read NetworkInterfaceNatRuleAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_nat_rule_association#update NetworkInterfaceNatRuleAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

