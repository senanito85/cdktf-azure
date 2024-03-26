package iothubendpointstoragecontainer


type IothubEndpointStorageContainerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_storage_container#create IothubEndpointStorageContainer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_storage_container#delete IothubEndpointStorageContainer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_storage_container#read IothubEndpointStorageContainer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub_endpoint_storage_container#update IothubEndpointStorageContainer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

