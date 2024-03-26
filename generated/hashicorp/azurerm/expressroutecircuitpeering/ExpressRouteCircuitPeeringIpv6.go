package expressroutecircuitpeering


type ExpressRouteCircuitPeeringIpv6 struct {
	// microsoft_peering block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#microsoft_peering ExpressRouteCircuitPeering#microsoft_peering}
	MicrosoftPeering *ExpressRouteCircuitPeeringIpv6MicrosoftPeering `field:"required" json:"microsoftPeering" yaml:"microsoftPeering"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#primary_peer_address_prefix ExpressRouteCircuitPeering#primary_peer_address_prefix}.
	PrimaryPeerAddressPrefix *string `field:"required" json:"primaryPeerAddressPrefix" yaml:"primaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#secondary_peer_address_prefix ExpressRouteCircuitPeering#secondary_peer_address_prefix}.
	SecondaryPeerAddressPrefix *string `field:"required" json:"secondaryPeerAddressPrefix" yaml:"secondaryPeerAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#route_filter_id ExpressRouteCircuitPeering#route_filter_id}.
	RouteFilterId *string `field:"optional" json:"routeFilterId" yaml:"routeFilterId"`
}

