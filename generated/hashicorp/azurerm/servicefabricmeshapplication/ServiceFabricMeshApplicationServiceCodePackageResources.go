package servicefabricmeshapplication


type ServiceFabricMeshApplicationServiceCodePackageResources struct {
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#requests ServiceFabricMeshApplication#requests}
	Requests *ServiceFabricMeshApplicationServiceCodePackageResourcesRequests `field:"required" json:"requests" yaml:"requests"`
	// limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#limits ServiceFabricMeshApplication#limits}
	Limits *ServiceFabricMeshApplicationServiceCodePackageResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
}

