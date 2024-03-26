package iothubendpointeventhub


type IothubEndpointEventhubTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_eventhub#create IothubEndpointEventhub#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_eventhub#delete IothubEndpointEventhub#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_eventhub#read IothubEndpointEventhub#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_eventhub#update IothubEndpointEventhub#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

