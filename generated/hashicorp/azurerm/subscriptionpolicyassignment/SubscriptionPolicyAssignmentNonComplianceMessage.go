package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentNonComplianceMessage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subscription_policy_assignment#content SubscriptionPolicyAssignment#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subscription_policy_assignment#policy_definition_reference_id SubscriptionPolicyAssignment#policy_definition_reference_id}.
	PolicyDefinitionReferenceId *string `field:"optional" json:"policyDefinitionReferenceId" yaml:"policyDefinitionReferenceId"`
}

