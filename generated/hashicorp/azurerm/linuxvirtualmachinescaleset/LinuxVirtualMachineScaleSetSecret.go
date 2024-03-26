package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/linux_virtual_machine_scale_set#certificate LinuxVirtualMachineScaleSet#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/linux_virtual_machine_scale_set#key_vault_id LinuxVirtualMachineScaleSet#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

