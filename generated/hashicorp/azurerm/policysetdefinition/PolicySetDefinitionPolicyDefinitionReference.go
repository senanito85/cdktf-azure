package policysetdefinition


type PolicySetDefinitionPolicyDefinitionReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_set_definition#policy_definition_id PolicySetDefinition#policy_definition_id}.
	PolicyDefinitionId *string `field:"required" json:"policyDefinitionId" yaml:"policyDefinitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_set_definition#parameters PolicySetDefinition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_set_definition#parameter_values PolicySetDefinition#parameter_values}.
	ParameterValues *string `field:"optional" json:"parameterValues" yaml:"parameterValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_set_definition#policy_group_names PolicySetDefinition#policy_group_names}.
	PolicyGroupNames *[]*string `field:"optional" json:"policyGroupNames" yaml:"policyGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/policy_set_definition#reference_id PolicySetDefinition#reference_id}.
	ReferenceId *string `field:"optional" json:"referenceId" yaml:"referenceId"`
}

