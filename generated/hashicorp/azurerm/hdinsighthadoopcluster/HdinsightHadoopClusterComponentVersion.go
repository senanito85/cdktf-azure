package hdinsighthadoopcluster


type HdinsightHadoopClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#hadoop HdinsightHadoopCluster#hadoop}.
	Hadoop *string `field:"required" json:"hadoop" yaml:"hadoop"`
}

