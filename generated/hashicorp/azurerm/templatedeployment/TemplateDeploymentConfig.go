package templatedeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TemplateDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#deployment_mode TemplateDeployment#deployment_mode}.
	DeploymentMode *string `field:"required" json:"deploymentMode" yaml:"deploymentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#name TemplateDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#resource_group_name TemplateDeployment#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#id TemplateDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#parameters TemplateDeployment#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#parameters_body TemplateDeployment#parameters_body}.
	ParametersBody *string `field:"optional" json:"parametersBody" yaml:"parametersBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#template_body TemplateDeployment#template_body}.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/template_deployment#timeouts TemplateDeployment#timeouts}
	Timeouts *TemplateDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

