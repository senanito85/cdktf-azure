package hdinsighthbasecluster


type HdinsightHbaseClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#head_node HdinsightHbaseCluster#head_node}
	HeadNode *HdinsightHbaseClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#worker_node HdinsightHbaseCluster#worker_node}
	WorkerNode *HdinsightHbaseClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#zookeeper_node HdinsightHbaseCluster#zookeeper_node}
	ZookeeperNode *HdinsightHbaseClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

