package vpngatewayconnection


type VpnGatewayConnectionRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_gateway_connection#associated_route_table VpnGatewayConnection#associated_route_table}.
	AssociatedRouteTable *string `field:"required" json:"associatedRouteTable" yaml:"associatedRouteTable"`
	// propagated_route_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_gateway_connection#propagated_route_table VpnGatewayConnection#propagated_route_table}
	PropagatedRouteTable *VpnGatewayConnectionRoutingPropagatedRouteTable `field:"optional" json:"propagatedRouteTable" yaml:"propagatedRouteTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/vpn_gateway_connection#propagated_route_tables VpnGatewayConnection#propagated_route_tables}.
	PropagatedRouteTables *[]*string `field:"optional" json:"propagatedRouteTables" yaml:"propagatedRouteTables"`
}

