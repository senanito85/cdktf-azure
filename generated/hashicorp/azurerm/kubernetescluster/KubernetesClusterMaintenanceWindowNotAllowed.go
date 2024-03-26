package kubernetescluster


type KubernetesClusterMaintenanceWindowNotAllowed struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#end KubernetesCluster#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#start KubernetesCluster#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

