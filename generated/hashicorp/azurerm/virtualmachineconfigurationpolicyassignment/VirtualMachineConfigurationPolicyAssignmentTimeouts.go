package virtualmachineconfigurationpolicyassignment


type VirtualMachineConfigurationPolicyAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#create VirtualMachineConfigurationPolicyAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#delete VirtualMachineConfigurationPolicyAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#read VirtualMachineConfigurationPolicyAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#update VirtualMachineConfigurationPolicyAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

