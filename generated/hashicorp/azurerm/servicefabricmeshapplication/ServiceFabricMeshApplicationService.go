package servicefabricmeshapplication


type ServiceFabricMeshApplicationService struct {
	// code_package block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#code_package ServiceFabricMeshApplication#code_package}
	CodePackage interface{} `field:"required" json:"codePackage" yaml:"codePackage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#name ServiceFabricMeshApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_application#os_type ServiceFabricMeshApplication#os_type}.
	OsType *string `field:"required" json:"osType" yaml:"osType"`
}

