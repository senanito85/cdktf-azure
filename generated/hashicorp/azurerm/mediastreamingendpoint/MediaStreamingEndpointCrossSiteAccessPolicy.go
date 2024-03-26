package mediastreamingendpoint


type MediaStreamingEndpointCrossSiteAccessPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_endpoint#client_access_policy MediaStreamingEndpoint#client_access_policy}.
	ClientAccessPolicy *string `field:"optional" json:"clientAccessPolicy" yaml:"clientAccessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_endpoint#cross_domain_policy MediaStreamingEndpoint#cross_domain_policy}.
	CrossDomainPolicy *string `field:"optional" json:"crossDomainPolicy" yaml:"crossDomainPolicy"`
}

