package hdinsightsparkcluster


type HdinsightSparkClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_spark_cluster#connection_direction HdinsightSparkCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_spark_cluster#private_link_enabled HdinsightSparkCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}

