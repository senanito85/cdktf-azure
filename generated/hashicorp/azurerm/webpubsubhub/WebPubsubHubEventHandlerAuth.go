package webpubsubhub


type WebPubsubHubEventHandlerAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_pubsub_hub#managed_identity_id WebPubsubHub#managed_identity_id}.
	ManagedIdentityId *string `field:"required" json:"managedIdentityId" yaml:"managedIdentityId"`
}

