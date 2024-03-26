package virtualnetworkgateway


type VirtualNetworkGatewayVpnClientConfigurationRootCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_gateway#name VirtualNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_network_gateway#public_cert_data VirtualNetworkGateway#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

