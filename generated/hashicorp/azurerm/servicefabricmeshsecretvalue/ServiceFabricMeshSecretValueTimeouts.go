package servicefabricmeshsecretvalue


type ServiceFabricMeshSecretValueTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_secret_value#create ServiceFabricMeshSecretValue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_secret_value#delete ServiceFabricMeshSecretValue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_secret_value#read ServiceFabricMeshSecretValue#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_mesh_secret_value#update ServiceFabricMeshSecretValue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

