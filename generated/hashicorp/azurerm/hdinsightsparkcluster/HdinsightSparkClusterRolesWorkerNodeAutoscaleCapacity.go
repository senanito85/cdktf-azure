package hdinsightsparkcluster


type HdinsightSparkClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_spark_cluster#max_instance_count HdinsightSparkCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_spark_cluster#min_instance_count HdinsightSparkCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

