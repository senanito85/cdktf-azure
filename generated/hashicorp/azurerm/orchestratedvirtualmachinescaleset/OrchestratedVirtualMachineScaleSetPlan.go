package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#name OrchestratedVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#product OrchestratedVirtualMachineScaleSet#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#publisher OrchestratedVirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

