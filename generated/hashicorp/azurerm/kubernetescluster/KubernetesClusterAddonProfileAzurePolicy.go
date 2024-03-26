package kubernetescluster


type KubernetesClusterAddonProfileAzurePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

