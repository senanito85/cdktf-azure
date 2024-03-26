package hdinsightrservercluster


type HdinsightRserverClusterRoles struct {
	// edge_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_rserver_cluster#edge_node HdinsightRserverCluster#edge_node}
	EdgeNode *HdinsightRserverClusterRolesEdgeNode `field:"required" json:"edgeNode" yaml:"edgeNode"`
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_rserver_cluster#head_node HdinsightRserverCluster#head_node}
	HeadNode *HdinsightRserverClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_rserver_cluster#worker_node HdinsightRserverCluster#worker_node}
	WorkerNode *HdinsightRserverClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_rserver_cluster#zookeeper_node HdinsightRserverCluster#zookeeper_node}
	ZookeeperNode *HdinsightRserverClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

