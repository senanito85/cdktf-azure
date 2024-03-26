package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/linux_virtual_machine_scale_set#type LinuxVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/linux_virtual_machine_scale_set#identity_ids LinuxVirtualMachineScaleSet#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

