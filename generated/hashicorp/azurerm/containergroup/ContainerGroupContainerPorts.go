package containergroup


type ContainerGroupContainerPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#port ContainerGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/container_group#protocol ContainerGroup#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

