package iottimeseriesinsightsaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotTimeSeriesInsightsAccessPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#name IotTimeSeriesInsightsAccessPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#principal_object_id IotTimeSeriesInsightsAccessPolicy#principal_object_id}.
	PrincipalObjectId *string `field:"required" json:"principalObjectId" yaml:"principalObjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#roles IotTimeSeriesInsightsAccessPolicy#roles}.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#time_series_insights_environment_id IotTimeSeriesInsightsAccessPolicy#time_series_insights_environment_id}.
	TimeSeriesInsightsEnvironmentId *string `field:"required" json:"timeSeriesInsightsEnvironmentId" yaml:"timeSeriesInsightsEnvironmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#description IotTimeSeriesInsightsAccessPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#id IotTimeSeriesInsightsAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_time_series_insights_access_policy#timeouts IotTimeSeriesInsightsAccessPolicy#timeouts}
	Timeouts *IotTimeSeriesInsightsAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

