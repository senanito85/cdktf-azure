package machinelearningcomputeinstance


type MachineLearningComputeInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_compute_instance#create MachineLearningComputeInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_compute_instance#delete MachineLearningComputeInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/machine_learning_compute_instance#read MachineLearningComputeInstance#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

