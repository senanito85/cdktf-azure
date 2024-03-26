package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetAutomaticInstanceRepair struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#enabled OrchestratedVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#grace_period OrchestratedVirtualMachineScaleSet#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

