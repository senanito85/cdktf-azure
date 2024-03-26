package datafactoryintegrationruntimeselfhosted

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryIntegrationRuntimeSelfHostedConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#name DataFactoryIntegrationRuntimeSelfHosted#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#resource_group_name DataFactoryIntegrationRuntimeSelfHosted#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#data_factory_id DataFactoryIntegrationRuntimeSelfHosted#data_factory_id}.
	DataFactoryId *string `field:"optional" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#data_factory_name DataFactoryIntegrationRuntimeSelfHosted#data_factory_name}.
	DataFactoryName *string `field:"optional" json:"dataFactoryName" yaml:"dataFactoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#description DataFactoryIntegrationRuntimeSelfHosted#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#id DataFactoryIntegrationRuntimeSelfHosted#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// rbac_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#rbac_authorization DataFactoryIntegrationRuntimeSelfHosted#rbac_authorization}
	RbacAuthorization interface{} `field:"optional" json:"rbacAuthorization" yaml:"rbacAuthorization"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory_integration_runtime_self_hosted#timeouts DataFactoryIntegrationRuntimeSelfHosted#timeouts}
	Timeouts *DataFactoryIntegrationRuntimeSelfHostedTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

