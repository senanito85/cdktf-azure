package hdinsighthadoopcluster


type HdinsightHadoopClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#capacity HdinsightHadoopCluster#capacity}
	Capacity *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#recurrence HdinsightHadoopCluster#recurrence}
	Recurrence *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

