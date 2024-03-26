package cdnendpoint


type CdnEndpointDeliveryRuleUrlRedirectAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#redirect_type CdnEndpoint#redirect_type}.
	RedirectType *string `field:"required" json:"redirectType" yaml:"redirectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#fragment CdnEndpoint#fragment}.
	Fragment *string `field:"optional" json:"fragment" yaml:"fragment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#hostname CdnEndpoint#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#path CdnEndpoint#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#protocol CdnEndpoint#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cdn_endpoint#query_string CdnEndpoint#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

