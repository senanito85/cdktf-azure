package dataazurermiothubdpssharedaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermIothubDpsSharedAccessPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/iothub_dps_shared_access_policy#iothub_dps_name DataAzurermIothubDpsSharedAccessPolicy#iothub_dps_name}.
	IothubDpsName *string `field:"required" json:"iothubDpsName" yaml:"iothubDpsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/iothub_dps_shared_access_policy#name DataAzurermIothubDpsSharedAccessPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/iothub_dps_shared_access_policy#resource_group_name DataAzurermIothubDpsSharedAccessPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/iothub_dps_shared_access_policy#id DataAzurermIothubDpsSharedAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/iothub_dps_shared_access_policy#timeouts DataAzurermIothubDpsSharedAccessPolicy#timeouts}
	Timeouts *DataAzurermIothubDpsSharedAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

