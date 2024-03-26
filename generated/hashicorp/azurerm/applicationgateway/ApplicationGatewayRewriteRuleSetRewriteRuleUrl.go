package applicationgateway


type ApplicationGatewayRewriteRuleSetRewriteRuleUrl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#path ApplicationGateway#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#query_string ApplicationGateway#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/application_gateway#reroute ApplicationGateway#reroute}.
	Reroute interface{} `field:"optional" json:"reroute" yaml:"reroute"`
}

