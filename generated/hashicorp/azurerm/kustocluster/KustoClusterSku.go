package kustocluster


type KustoClusterSku struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kusto_cluster#name KustoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kusto_cluster#capacity KustoCluster#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

