package windowsvirtualmachine


type WindowsVirtualMachinePlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine#name WindowsVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine#product WindowsVirtualMachine#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine#publisher WindowsVirtualMachine#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}

