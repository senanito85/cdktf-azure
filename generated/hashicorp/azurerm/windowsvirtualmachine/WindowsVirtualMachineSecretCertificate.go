package windowsvirtualmachine


type WindowsVirtualMachineSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine#store WindowsVirtualMachine#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/windows_virtual_machine#url WindowsVirtualMachine#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

