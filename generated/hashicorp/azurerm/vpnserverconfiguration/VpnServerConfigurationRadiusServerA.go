package vpnserverconfiguration


type VpnServerConfigurationRadiusServerA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#address VpnServerConfiguration#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#secret VpnServerConfiguration#secret}.
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// server_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#server_root_certificate VpnServerConfiguration#server_root_certificate}
	ServerRootCertificate interface{} `field:"required" json:"serverRootCertificate" yaml:"serverRootCertificate"`
	// client_root_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#client_root_certificate VpnServerConfiguration#client_root_certificate}
	ClientRootCertificate interface{} `field:"optional" json:"clientRootCertificate" yaml:"clientRootCertificate"`
}

