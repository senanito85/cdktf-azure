package streamanalyticsoutputmssql

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StreamAnalyticsOutputMssqlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#database StreamAnalyticsOutputMssql#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#name StreamAnalyticsOutputMssql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#password StreamAnalyticsOutputMssql#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#resource_group_name StreamAnalyticsOutputMssql#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#server StreamAnalyticsOutputMssql#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#stream_analytics_job_name StreamAnalyticsOutputMssql#stream_analytics_job_name}.
	StreamAnalyticsJobName *string `field:"required" json:"streamAnalyticsJobName" yaml:"streamAnalyticsJobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#table StreamAnalyticsOutputMssql#table}.
	Table *string `field:"required" json:"table" yaml:"table"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#user StreamAnalyticsOutputMssql#user}.
	User *string `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#id StreamAnalyticsOutputMssql#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/stream_analytics_output_mssql#timeouts StreamAnalyticsOutputMssql#timeouts}
	Timeouts *StreamAnalyticsOutputMssqlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

