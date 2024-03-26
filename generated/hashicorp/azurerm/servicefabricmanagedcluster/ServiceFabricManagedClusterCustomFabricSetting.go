package servicefabricmanagedcluster


type ServiceFabricManagedClusterCustomFabricSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_managed_cluster#parameter ServiceFabricManagedCluster#parameter}.
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_managed_cluster#section ServiceFabricManagedCluster#section}.
	Section *string `field:"required" json:"section" yaml:"section"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_managed_cluster#value ServiceFabricManagedCluster#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

