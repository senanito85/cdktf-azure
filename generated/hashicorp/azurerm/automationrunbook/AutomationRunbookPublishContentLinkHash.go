package automationrunbook


type AutomationRunbookPublishContentLinkHash struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/automation_runbook#algorithm AutomationRunbook#algorithm}.
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/automation_runbook#value AutomationRunbook#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

