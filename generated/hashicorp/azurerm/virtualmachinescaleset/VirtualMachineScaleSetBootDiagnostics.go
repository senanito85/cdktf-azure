package virtualmachinescaleset


type VirtualMachineScaleSetBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#storage_uri VirtualMachineScaleSet#storage_uri}.
	StorageUri *string `field:"required" json:"storageUri" yaml:"storageUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine_scale_set#enabled VirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

