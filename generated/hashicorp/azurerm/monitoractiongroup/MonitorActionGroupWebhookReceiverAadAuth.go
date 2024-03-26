package monitoractiongroup


type MonitorActionGroupWebhookReceiverAadAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_action_group#object_id MonitorActionGroup#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_action_group#identifier_uri MonitorActionGroup#identifier_uri}.
	IdentifierUri *string `field:"optional" json:"identifierUri" yaml:"identifierUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/monitor_action_group#tenant_id MonitorActionGroup#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

