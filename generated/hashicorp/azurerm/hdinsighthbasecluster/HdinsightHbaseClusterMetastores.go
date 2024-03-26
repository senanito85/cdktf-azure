package hdinsighthbasecluster


type HdinsightHbaseClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#ambari HdinsightHbaseCluster#ambari}
	Ambari *HdinsightHbaseClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#hive HdinsightHbaseCluster#hive}
	Hive *HdinsightHbaseClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#oozie HdinsightHbaseCluster#oozie}
	Oozie *HdinsightHbaseClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

