package datafactorydatasetjson


type DataFactoryDatasetJsonSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_json#name DataFactoryDatasetJson#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_json#description DataFactoryDatasetJson#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_json#type DataFactoryDatasetJson#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

