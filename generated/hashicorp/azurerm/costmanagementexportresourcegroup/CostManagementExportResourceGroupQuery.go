package costmanagementexportresourcegroup


type CostManagementExportResourceGroupQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#time_frame CostManagementExportResourceGroup#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cost_management_export_resource_group#type CostManagementExportResourceGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

