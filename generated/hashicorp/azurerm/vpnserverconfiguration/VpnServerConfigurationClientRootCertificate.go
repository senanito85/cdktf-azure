package vpnserverconfiguration


type VpnServerConfigurationClientRootCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#name VpnServerConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#public_cert_data VpnServerConfiguration#public_cert_data}.
	PublicCertData *string `field:"required" json:"publicCertData" yaml:"publicCertData"`
}

