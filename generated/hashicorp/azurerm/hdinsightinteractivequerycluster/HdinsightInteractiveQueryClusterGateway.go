package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterGateway struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_interactive_query_cluster#enabled HdinsightInteractiveQueryCluster#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

