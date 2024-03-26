package servicefabriccluster


type ServiceFabricClusterUpgradePolicyHealthPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_cluster#max_unhealthy_applications_percent ServiceFabricCluster#max_unhealthy_applications_percent}.
	MaxUnhealthyApplicationsPercent *float64 `field:"optional" json:"maxUnhealthyApplicationsPercent" yaml:"maxUnhealthyApplicationsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/service_fabric_cluster#max_unhealthy_nodes_percent ServiceFabricCluster#max_unhealthy_nodes_percent}.
	MaxUnhealthyNodesPercent *float64 `field:"optional" json:"maxUnhealthyNodesPercent" yaml:"maxUnhealthyNodesPercent"`
}

