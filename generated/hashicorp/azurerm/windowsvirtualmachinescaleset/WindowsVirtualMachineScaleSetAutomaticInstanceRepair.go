package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetAutomaticInstanceRepair struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine_scale_set#grace_period WindowsVirtualMachineScaleSet#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

