package kubernetescluster


type KubernetesClusterHttpProxyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#http_proxy KubernetesCluster#http_proxy}.
	HttpProxy *string `field:"optional" json:"httpProxy" yaml:"httpProxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#https_proxy KubernetesCluster#https_proxy}.
	HttpsProxy *string `field:"optional" json:"httpsProxy" yaml:"httpsProxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#no_proxy KubernetesCluster#no_proxy}.
	NoProxy *[]*string `field:"optional" json:"noProxy" yaml:"noProxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#trusted_ca KubernetesCluster#trusted_ca}.
	TrustedCa *string `field:"optional" json:"trustedCa" yaml:"trustedCa"`
}

