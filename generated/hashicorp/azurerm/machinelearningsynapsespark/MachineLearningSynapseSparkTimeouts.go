package machinelearningsynapsespark


type MachineLearningSynapseSparkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_synapse_spark#create MachineLearningSynapseSpark#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_synapse_spark#delete MachineLearningSynapseSpark#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_synapse_spark#read MachineLearningSynapseSpark#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

