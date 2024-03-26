package eventgrideventsubscription


type EventgridEventSubscriptionStorageBlobDeadLetterDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_event_subscription#storage_account_id EventgridEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventgrid_event_subscription#storage_blob_container_name EventgridEventSubscription#storage_blob_container_name}.
	StorageBlobContainerName *string `field:"required" json:"storageBlobContainerName" yaml:"storageBlobContainerName"`
}

