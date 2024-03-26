package servicefabricmanagedcluster


type ServiceFabricManagedClusterNodeTypeVmSecretsCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_managed_cluster#store ServiceFabricManagedCluster#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_managed_cluster#url ServiceFabricManagedCluster#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

