package hdinsighthadoopcluster


type HdinsightHadoopClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#filesystem_id HdinsightHadoopCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#is_default HdinsightHadoopCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#managed_identity_resource_id HdinsightHadoopCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hadoop_cluster#storage_resource_id HdinsightHadoopCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

