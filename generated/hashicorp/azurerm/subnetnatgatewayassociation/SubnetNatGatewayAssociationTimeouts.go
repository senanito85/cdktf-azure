package subnetnatgatewayassociation


type SubnetNatGatewayAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subnet_nat_gateway_association#create SubnetNatGatewayAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subnet_nat_gateway_association#delete SubnetNatGatewayAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subnet_nat_gateway_association#read SubnetNatGatewayAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subnet_nat_gateway_association#update SubnetNatGatewayAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

