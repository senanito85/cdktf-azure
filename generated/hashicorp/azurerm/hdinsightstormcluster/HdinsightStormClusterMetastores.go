package hdinsightstormcluster


type HdinsightStormClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_storm_cluster#ambari HdinsightStormCluster#ambari}
	Ambari *HdinsightStormClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_storm_cluster#hive HdinsightStormCluster#hive}
	Hive *HdinsightStormClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_storm_cluster#oozie HdinsightStormCluster#oozie}
	Oozie *HdinsightStormClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

