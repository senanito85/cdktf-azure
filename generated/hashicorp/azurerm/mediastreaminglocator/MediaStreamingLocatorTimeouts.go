package mediastreaminglocator


type MediaStreamingLocatorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_locator#create MediaStreamingLocator#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_locator#delete MediaStreamingLocator#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_locator#read MediaStreamingLocator#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

