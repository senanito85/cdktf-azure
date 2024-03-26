package hdinsightinteractivequerycluster


type HdinsightInteractiveQueryClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_interactive_query_cluster#connection_direction HdinsightInteractiveQueryCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_interactive_query_cluster#private_link_enabled HdinsightInteractiveQueryCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}

