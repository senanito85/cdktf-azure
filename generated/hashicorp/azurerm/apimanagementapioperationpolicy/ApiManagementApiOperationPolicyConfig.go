package apimanagementapioperationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementApiOperationPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#api_management_name ApiManagementApiOperationPolicy#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#api_name ApiManagementApiOperationPolicy#api_name}.
	ApiName *string `field:"required" json:"apiName" yaml:"apiName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#operation_id ApiManagementApiOperationPolicy#operation_id}.
	OperationId *string `field:"required" json:"operationId" yaml:"operationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#resource_group_name ApiManagementApiOperationPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#id ApiManagementApiOperationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#timeouts ApiManagementApiOperationPolicy#timeouts}
	Timeouts *ApiManagementApiOperationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#xml_content ApiManagementApiOperationPolicy#xml_content}.
	XmlContent *string `field:"optional" json:"xmlContent" yaml:"xmlContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/api_management_api_operation_policy#xml_link ApiManagementApiOperationPolicy#xml_link}.
	XmlLink *string `field:"optional" json:"xmlLink" yaml:"xmlLink"`
}

