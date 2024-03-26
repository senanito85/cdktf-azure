package virtualhub


type VirtualHubRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_hub#address_prefixes VirtualHub#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_hub#next_hop_ip_address VirtualHub#next_hop_ip_address}.
	NextHopIpAddress *string `field:"required" json:"nextHopIpAddress" yaml:"nextHopIpAddress"`
}

