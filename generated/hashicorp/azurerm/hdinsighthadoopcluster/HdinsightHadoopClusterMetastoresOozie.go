package hdinsighthadoopcluster


type HdinsightHadoopClusterMetastoresOozie struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#database_name HdinsightHadoopCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#server HdinsightHadoopCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

