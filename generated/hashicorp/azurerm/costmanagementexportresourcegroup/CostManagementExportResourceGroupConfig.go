package costmanagementexportresourcegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CostManagementExportResourceGroupConfig struct {
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
	// delivery_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#delivery_info CostManagementExportResourceGroup#delivery_info}
	DeliveryInfo *CostManagementExportResourceGroupDeliveryInfo `field:"required" json:"deliveryInfo" yaml:"deliveryInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#name CostManagementExportResourceGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#query CostManagementExportResourceGroup#query}
	Query *CostManagementExportResourceGroupQuery `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#recurrence_period_end CostManagementExportResourceGroup#recurrence_period_end}.
	RecurrencePeriodEnd *string `field:"required" json:"recurrencePeriodEnd" yaml:"recurrencePeriodEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#recurrence_period_start CostManagementExportResourceGroup#recurrence_period_start}.
	RecurrencePeriodStart *string `field:"required" json:"recurrencePeriodStart" yaml:"recurrencePeriodStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#recurrence_type CostManagementExportResourceGroup#recurrence_type}.
	RecurrenceType *string `field:"required" json:"recurrenceType" yaml:"recurrenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#resource_group_id CostManagementExportResourceGroup#resource_group_id}.
	ResourceGroupId *string `field:"required" json:"resourceGroupId" yaml:"resourceGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#active CostManagementExportResourceGroup#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#id CostManagementExportResourceGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#timeouts CostManagementExportResourceGroup#timeouts}
	Timeouts *CostManagementExportResourceGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

