package kustocluster


type KustoClusterOptimizedAutoScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kusto_cluster#maximum_instances KustoCluster#maximum_instances}.
	MaximumInstances *float64 `field:"required" json:"maximumInstances" yaml:"maximumInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/kusto_cluster#minimum_instances KustoCluster#minimum_instances}.
	MinimumInstances *float64 `field:"required" json:"minimumInstances" yaml:"minimumInstances"`
}

