package subscriptioncostmanagementexport


type SubscriptionCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subscription_cost_management_export#time_frame SubscriptionCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/subscription_cost_management_export#type SubscriptionCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

