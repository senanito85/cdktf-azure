package virtualmachineconfigurationpolicyassignment


type VirtualMachineConfigurationPolicyAssignmentConfigurationParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#name VirtualMachineConfigurationPolicyAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#value VirtualMachineConfigurationPolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

