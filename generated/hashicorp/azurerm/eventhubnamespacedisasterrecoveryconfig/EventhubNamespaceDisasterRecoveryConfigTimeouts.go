package eventhubnamespacedisasterrecoveryconfig


type EventhubNamespaceDisasterRecoveryConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace_disaster_recovery_config#create EventhubNamespaceDisasterRecoveryConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace_disaster_recovery_config#delete EventhubNamespaceDisasterRecoveryConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace_disaster_recovery_config#read EventhubNamespaceDisasterRecoveryConfig#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub_namespace_disaster_recovery_config#update EventhubNamespaceDisasterRecoveryConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

