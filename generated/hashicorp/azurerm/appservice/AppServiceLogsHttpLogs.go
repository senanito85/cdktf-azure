package appservice


type AppServiceLogsHttpLogs struct {
	// azure_blob_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service#azure_blob_storage AppService#azure_blob_storage}
	AzureBlobStorage *AppServiceLogsHttpLogsAzureBlobStorage `field:"optional" json:"azureBlobStorage" yaml:"azureBlobStorage"`
	// file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service#file_system AppService#file_system}
	FileSystem *AppServiceLogsHttpLogsFileSystem `field:"optional" json:"fileSystem" yaml:"fileSystem"`
}

