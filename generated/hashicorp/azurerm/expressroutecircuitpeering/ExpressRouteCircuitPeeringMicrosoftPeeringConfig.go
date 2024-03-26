package expressroutecircuitpeering


type ExpressRouteCircuitPeeringMicrosoftPeeringConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#advertised_public_prefixes ExpressRouteCircuitPeering#advertised_public_prefixes}.
	AdvertisedPublicPrefixes *[]*string `field:"required" json:"advertisedPublicPrefixes" yaml:"advertisedPublicPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#customer_asn ExpressRouteCircuitPeering#customer_asn}.
	CustomerAsn *float64 `field:"optional" json:"customerAsn" yaml:"customerAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_circuit_peering#routing_registry_name ExpressRouteCircuitPeering#routing_registry_name}.
	RoutingRegistryName *string `field:"optional" json:"routingRegistryName" yaml:"routingRegistryName"`
}

