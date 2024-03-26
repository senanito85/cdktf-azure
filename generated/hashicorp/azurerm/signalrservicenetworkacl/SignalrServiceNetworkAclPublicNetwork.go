package signalrservicenetworkacl


type SignalrServiceNetworkAclPublicNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service_network_acl#allowed_request_types SignalrServiceNetworkAcl#allowed_request_types}.
	AllowedRequestTypes *[]*string `field:"optional" json:"allowedRequestTypes" yaml:"allowedRequestTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/signalr_service_network_acl#denied_request_types SignalrServiceNetworkAcl#denied_request_types}.
	DeniedRequestTypes *[]*string `field:"optional" json:"deniedRequestTypes" yaml:"deniedRequestTypes"`
}

