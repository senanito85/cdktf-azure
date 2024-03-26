package hdinsighthbasecluster


type HdinsightHbaseClusterMonitor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#log_analytics_workspace_id HdinsightHbaseCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_hbase_cluster#primary_key HdinsightHbaseCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

