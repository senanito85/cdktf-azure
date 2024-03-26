package virtualmachinescaleset


type VirtualMachineScaleSetPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#product VirtualMachineScaleSet#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#publisher VirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

