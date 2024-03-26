package kubernetescluster


type KubernetesClusterMaintenanceWindowAllowed struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#hours KubernetesCluster#hours}.
	Hours *[]*float64 `field:"required" json:"hours" yaml:"hours"`
}

