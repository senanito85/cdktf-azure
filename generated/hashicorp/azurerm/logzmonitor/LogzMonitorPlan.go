package logzmonitor


type LogzMonitorPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logz_monitor#billing_cycle LogzMonitor#billing_cycle}.
	BillingCycle *string `field:"required" json:"billingCycle" yaml:"billingCycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logz_monitor#effective_date LogzMonitor#effective_date}.
	EffectiveDate *string `field:"required" json:"effectiveDate" yaml:"effectiveDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logz_monitor#plan_id LogzMonitor#plan_id}.
	PlanId *string `field:"required" json:"planId" yaml:"planId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logz_monitor#usage_type LogzMonitor#usage_type}.
	UsageType *string `field:"required" json:"usageType" yaml:"usageType"`
}

