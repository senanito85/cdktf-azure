package datalakestore


type DataLakeStoreIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_lake_store#type DataLakeStore#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

