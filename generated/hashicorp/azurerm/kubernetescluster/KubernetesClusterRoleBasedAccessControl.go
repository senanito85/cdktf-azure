package kubernetescluster


type KubernetesClusterRoleBasedAccessControl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// azure_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#azure_active_directory KubernetesCluster#azure_active_directory}
	AzureActiveDirectory *KubernetesClusterRoleBasedAccessControlAzureActiveDirectory `field:"optional" json:"azureActiveDirectory" yaml:"azureActiveDirectory"`
}

