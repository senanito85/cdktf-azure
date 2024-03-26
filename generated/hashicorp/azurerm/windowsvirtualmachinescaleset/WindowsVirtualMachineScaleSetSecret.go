package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine_scale_set#certificate WindowsVirtualMachineScaleSet#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine_scale_set#key_vault_id WindowsVirtualMachineScaleSet#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}

