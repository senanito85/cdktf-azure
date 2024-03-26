package containergroup


type ContainerGroupContainerLivenessProbeHttpGet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#path ContainerGroup#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#port ContainerGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#scheme ContainerGroup#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

