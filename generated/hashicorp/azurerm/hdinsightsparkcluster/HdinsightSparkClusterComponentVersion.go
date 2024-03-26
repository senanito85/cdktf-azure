package hdinsightsparkcluster


type HdinsightSparkClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_spark_cluster#spark HdinsightSparkCluster#spark}.
	Spark *string `field:"required" json:"spark" yaml:"spark"`
}

