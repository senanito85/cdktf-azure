package kubernetescluster


type KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#enabled KubernetesCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#secret_rotation_enabled KubernetesCluster#secret_rotation_enabled}.
	SecretRotationEnabled interface{} `field:"optional" json:"secretRotationEnabled" yaml:"secretRotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#secret_rotation_interval KubernetesCluster#secret_rotation_interval}.
	SecretRotationInterval *string `field:"optional" json:"secretRotationInterval" yaml:"secretRotationInterval"`
}

