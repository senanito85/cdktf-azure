package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#public_key OrchestratedVirtualMachineScaleSet#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#username OrchestratedVirtualMachineScaleSet#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

