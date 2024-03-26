package hdinsightkafkacluster


type HdinsightKafkaClusterRestProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_kafka_cluster#security_group_id HdinsightKafkaCluster#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hdinsight_kafka_cluster#security_group_name HdinsightKafkaCluster#security_group_name}.
	SecurityGroupName *string `field:"optional" json:"securityGroupName" yaml:"securityGroupName"`
}

