package batchpool


type BatchPoolContainerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#container_image_names BatchPool#container_image_names}.
	ContainerImageNames *[]*string `field:"optional" json:"containerImageNames" yaml:"containerImageNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#container_registries BatchPool#container_registries}.
	ContainerRegistries interface{} `field:"optional" json:"containerRegistries" yaml:"containerRegistries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#type BatchPool#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

