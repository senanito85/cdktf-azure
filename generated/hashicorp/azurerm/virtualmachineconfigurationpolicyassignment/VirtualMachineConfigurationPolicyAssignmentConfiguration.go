package virtualmachineconfigurationpolicyassignment


type VirtualMachineConfigurationPolicyAssignmentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#name VirtualMachineConfigurationPolicyAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#parameter VirtualMachineConfigurationPolicyAssignment#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_configuration_policy_assignment#version VirtualMachineConfigurationPolicyAssignment#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

