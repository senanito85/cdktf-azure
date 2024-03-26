package dataazurermbillingmpaaccountscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermBillingMpaAccountScopeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/billing_mpa_account_scope#billing_account_name DataAzurermBillingMpaAccountScope#billing_account_name}.
	BillingAccountName *string `field:"required" json:"billingAccountName" yaml:"billingAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/billing_mpa_account_scope#customer_name DataAzurermBillingMpaAccountScope#customer_name}.
	CustomerName *string `field:"required" json:"customerName" yaml:"customerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/billing_mpa_account_scope#id DataAzurermBillingMpaAccountScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/data-sources/billing_mpa_account_scope#timeouts DataAzurermBillingMpaAccountScope#timeouts}
	Timeouts *DataAzurermBillingMpaAccountScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

