package kubernetescluster


type KubernetesClusterAddonProfile struct {
	// aci_connector_linux block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#aci_connector_linux KubernetesCluster#aci_connector_linux}
	AciConnectorLinux *KubernetesClusterAddonProfileAciConnectorLinux `field:"optional" json:"aciConnectorLinux" yaml:"aciConnectorLinux"`
	// azure_keyvault_secrets_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#azure_keyvault_secrets_provider KubernetesCluster#azure_keyvault_secrets_provider}
	AzureKeyvaultSecretsProvider *KubernetesClusterAddonProfileAzureKeyvaultSecretsProvider `field:"optional" json:"azureKeyvaultSecretsProvider" yaml:"azureKeyvaultSecretsProvider"`
	// azure_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#azure_policy KubernetesCluster#azure_policy}
	AzurePolicy *KubernetesClusterAddonProfileAzurePolicy `field:"optional" json:"azurePolicy" yaml:"azurePolicy"`
	// http_application_routing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#http_application_routing KubernetesCluster#http_application_routing}
	HttpApplicationRouting *KubernetesClusterAddonProfileHttpApplicationRouting `field:"optional" json:"httpApplicationRouting" yaml:"httpApplicationRouting"`
	// ingress_application_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#ingress_application_gateway KubernetesCluster#ingress_application_gateway}
	IngressApplicationGateway *KubernetesClusterAddonProfileIngressApplicationGateway `field:"optional" json:"ingressApplicationGateway" yaml:"ingressApplicationGateway"`
	// kube_dashboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#kube_dashboard KubernetesCluster#kube_dashboard}
	KubeDashboard *KubernetesClusterAddonProfileKubeDashboard `field:"optional" json:"kubeDashboard" yaml:"kubeDashboard"`
	// oms_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#oms_agent KubernetesCluster#oms_agent}
	OmsAgent *KubernetesClusterAddonProfileOmsAgent `field:"optional" json:"omsAgent" yaml:"omsAgent"`
	// open_service_mesh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kubernetes_cluster#open_service_mesh KubernetesCluster#open_service_mesh}
	OpenServiceMesh *KubernetesClusterAddonProfileOpenServiceMesh `field:"optional" json:"openServiceMesh" yaml:"openServiceMesh"`
}

