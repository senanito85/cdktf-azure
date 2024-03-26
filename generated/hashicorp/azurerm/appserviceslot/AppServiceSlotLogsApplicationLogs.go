package appserviceslot


type AppServiceSlotLogsApplicationLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#azure_blob_storage AppServiceSlot#azure_blob_storage}
	AzureBlobStorage *AppServiceSlotLogsApplicationLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#file_system_level AppServiceSlot#file_system_level}.
	FileSystemLevel *string `field:"optional" json:"fileSystemLevel" yaml:"fileSystemLevel"`
}

