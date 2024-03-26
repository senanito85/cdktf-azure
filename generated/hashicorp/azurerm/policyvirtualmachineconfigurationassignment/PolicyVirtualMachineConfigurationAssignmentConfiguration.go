package policyvirtualmachineconfigurationassignment


type PolicyVirtualMachineConfigurationAssignmentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#assignment_type PolicyVirtualMachineConfigurationAssignment#assignment_type}.
	AssignmentType *string `field:"optional" json:"assignmentType" yaml:"assignmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#content_hash PolicyVirtualMachineConfigurationAssignment#content_hash}.
	ContentHash *string `field:"optional" json:"contentHash" yaml:"contentHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#content_uri PolicyVirtualMachineConfigurationAssignment#content_uri}.
	ContentUri *string `field:"optional" json:"contentUri" yaml:"contentUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#name PolicyVirtualMachineConfigurationAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#parameter PolicyVirtualMachineConfigurationAssignment#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_virtual_machine_configuration_assignment#version PolicyVirtualMachineConfigurationAssignment#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

