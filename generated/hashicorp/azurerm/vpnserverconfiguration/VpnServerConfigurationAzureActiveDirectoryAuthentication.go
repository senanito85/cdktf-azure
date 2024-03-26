package vpnserverconfiguration


type VpnServerConfigurationAzureActiveDirectoryAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#audience VpnServerConfiguration#audience}.
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#issuer VpnServerConfiguration#issuer}.
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_server_configuration#tenant VpnServerConfiguration#tenant}.
	Tenant *string `field:"required" json:"tenant" yaml:"tenant"`
}

