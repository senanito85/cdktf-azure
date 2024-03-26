package virtualnetworkdnsservers


type VirtualNetworkDnsServersTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_dns_servers#create VirtualNetworkDnsServers#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_dns_servers#delete VirtualNetworkDnsServers#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_dns_servers#read VirtualNetworkDnsServers#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_dns_servers#update VirtualNetworkDnsServers#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

