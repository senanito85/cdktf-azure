package mediajob


type MediaJobOutputAsset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_job#name MediaJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_job#label MediaJob#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
}

