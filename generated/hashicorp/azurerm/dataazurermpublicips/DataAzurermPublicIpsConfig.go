package dataazurermpublicips

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermPublicIpsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#resource_group_name DataAzurermPublicIps#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#allocation_type DataAzurermPublicIps#allocation_type}.
	AllocationType *string `field:"optional" json:"allocationType" yaml:"allocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#attached DataAzurermPublicIps#attached}.
	Attached interface{} `field:"optional" json:"attached" yaml:"attached"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#attachment_status DataAzurermPublicIps#attachment_status}.
	AttachmentStatus *string `field:"optional" json:"attachmentStatus" yaml:"attachmentStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#id DataAzurermPublicIps#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#name_prefix DataAzurermPublicIps#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/public_ips#timeouts DataAzurermPublicIps#timeouts}
	Timeouts *DataAzurermPublicIpsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

