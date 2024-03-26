package servicefabricmeshapplication


type ServiceFabricMeshApplicationServiceCodePackageResourcesRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#cpu ServiceFabricMeshApplication#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#memory ServiceFabricMeshApplication#memory}.
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
}

