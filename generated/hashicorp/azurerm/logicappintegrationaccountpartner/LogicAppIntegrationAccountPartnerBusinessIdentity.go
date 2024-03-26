package logicappintegrationaccountpartner


type LogicAppIntegrationAccountPartnerBusinessIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logic_app_integration_account_partner#qualifier LogicAppIntegrationAccountPartner#qualifier}.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logic_app_integration_account_partner#value LogicAppIntegrationAccountPartner#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

