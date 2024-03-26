package batchpool


type BatchPoolStartTaskResourceFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#auto_storage_container_name BatchPool#auto_storage_container_name}.
	AutoStorageContainerName *string `field:"optional" json:"autoStorageContainerName" yaml:"autoStorageContainerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#blob_prefix BatchPool#blob_prefix}.
	BlobPrefix *string `field:"optional" json:"blobPrefix" yaml:"blobPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#file_mode BatchPool#file_mode}.
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#file_path BatchPool#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#http_url BatchPool#http_url}.
	HttpUrl *string `field:"optional" json:"httpUrl" yaml:"httpUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#storage_container_url BatchPool#storage_container_url}.
	StorageContainerUrl *string `field:"optional" json:"storageContainerUrl" yaml:"storageContainerUrl"`
}

