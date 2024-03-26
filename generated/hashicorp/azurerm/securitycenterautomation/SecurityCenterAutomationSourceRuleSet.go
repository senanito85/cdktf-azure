package securitycenterautomation


type SecurityCenterAutomationSourceRuleSet struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/security_center_automation#rule SecurityCenterAutomation#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

