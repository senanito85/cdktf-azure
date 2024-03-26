package vpngateway


type VpnGatewayBgpSettingsInstance1BgpPeeringAddress struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_gateway#custom_ips VpnGateway#custom_ips}.
	CustomIps *[]*string `field:"required" json:"customIps" yaml:"customIps"`
}

