package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupTimePeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/consumption_budget_management_group#start_date ConsumptionBudgetManagementGroup#start_date}.
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/consumption_budget_management_group#end_date ConsumptionBudgetManagementGroup#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
}

