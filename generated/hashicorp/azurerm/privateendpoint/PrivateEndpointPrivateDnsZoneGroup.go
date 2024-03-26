package privateendpoint


type PrivateEndpointPrivateDnsZoneGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/private_endpoint#name PrivateEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/private_endpoint#private_dns_zone_ids PrivateEndpoint#private_dns_zone_ids}.
	PrivateDnsZoneIds *[]*string `field:"required" json:"privateDnsZoneIds" yaml:"privateDnsZoneIds"`
}

