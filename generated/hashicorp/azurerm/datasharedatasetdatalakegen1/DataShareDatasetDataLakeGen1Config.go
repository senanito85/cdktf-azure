package datasharedatasetdatalakegen1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataShareDatasetDataLakeGen1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#data_lake_store_id DataShareDatasetDataLakeGen1#data_lake_store_id}.
	DataLakeStoreId *string `field:"required" json:"dataLakeStoreId" yaml:"dataLakeStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#data_share_id DataShareDatasetDataLakeGen1#data_share_id}.
	DataShareId *string `field:"required" json:"dataShareId" yaml:"dataShareId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#folder_path DataShareDatasetDataLakeGen1#folder_path}.
	FolderPath *string `field:"required" json:"folderPath" yaml:"folderPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#name DataShareDatasetDataLakeGen1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#file_name DataShareDatasetDataLakeGen1#file_name}.
	FileName *string `field:"optional" json:"fileName" yaml:"fileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#id DataShareDatasetDataLakeGen1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_share_dataset_data_lake_gen1#timeouts DataShareDatasetDataLakeGen1#timeouts}
	Timeouts *DataShareDatasetDataLakeGen1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

