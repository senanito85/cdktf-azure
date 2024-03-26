package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationWinrmListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#protocol OrchestratedVirtualMachineScaleSet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/orchestrated_virtual_machine_scale_set#certificate_url OrchestratedVirtualMachineScaleSet#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}

