package hdinsightmlservicescluster


type HdinsightMlServicesClusterRoles struct {
	// edge_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_ml_services_cluster#edge_node HdinsightMlServicesCluster#edge_node}
	EdgeNode *HdinsightMlServicesClusterRolesEdgeNode `field:"required" json:"edgeNode" yaml:"edgeNode"`
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_ml_services_cluster#head_node HdinsightMlServicesCluster#head_node}
	HeadNode *HdinsightMlServicesClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_ml_services_cluster#worker_node HdinsightMlServicesCluster#worker_node}
	WorkerNode *HdinsightMlServicesClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_ml_services_cluster#zookeeper_node HdinsightMlServicesCluster#zookeeper_node}
	ZookeeperNode *HdinsightMlServicesClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

