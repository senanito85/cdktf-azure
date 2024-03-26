package kubernetescluster


type KubernetesClusterAddonProfileAciConnectorLinux struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#subnet_name KubernetesCluster#subnet_name}.
	SubnetName *string `field:"optional" json:"subnetName" yaml:"subnetName"`
}

