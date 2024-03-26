package servicefabricmeshapplication


type ServiceFabricMeshApplicationServiceCodePackage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#image_name ServiceFabricMeshApplication#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#name ServiceFabricMeshApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#resources ServiceFabricMeshApplication#resources}
	Resources *ServiceFabricMeshApplicationServiceCodePackageResources `field:"required" json:"resources" yaml:"resources"`
}

