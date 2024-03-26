package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileLinuxConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#disable_password_authentication VirtualMachineScaleSet#disable_password_authentication}.
	DisablePasswordAuthentication interface{} `field:"optional" json:"disablePasswordAuthentication" yaml:"disablePasswordAuthentication"`
	// ssh_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#ssh_keys VirtualMachineScaleSet#ssh_keys}
	SshKeys interface{} `field:"optional" json:"sshKeys" yaml:"sshKeys"`
}

