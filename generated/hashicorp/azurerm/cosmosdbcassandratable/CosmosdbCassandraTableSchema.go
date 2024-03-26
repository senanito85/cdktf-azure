package cosmosdbcassandratable


type CosmosdbCassandraTableSchema struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_cassandra_table#column CosmosdbCassandraTable#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// partition_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_cassandra_table#partition_key CosmosdbCassandraTable#partition_key}
	PartitionKey interface{} `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// cluster_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_cassandra_table#cluster_key CosmosdbCassandraTable#cluster_key}
	ClusterKey interface{} `field:"optional" json:"clusterKey" yaml:"clusterKey"`
}

