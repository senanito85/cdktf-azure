package datafactorydatasetsnowflake


type DataFactoryDatasetSnowflakeStructureColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_snowflake#name DataFactoryDatasetSnowflake#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_snowflake#description DataFactoryDatasetSnowflake#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_dataset_snowflake#type DataFactoryDatasetSnowflake#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

