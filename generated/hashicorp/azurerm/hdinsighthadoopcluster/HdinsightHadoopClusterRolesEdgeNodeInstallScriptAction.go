package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesEdgeNodeInstallScriptAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#uri HdinsightHadoopCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

