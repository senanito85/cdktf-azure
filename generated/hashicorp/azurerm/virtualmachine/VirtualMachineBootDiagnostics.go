package virtualmachine


type VirtualMachineBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#enabled VirtualMachine#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#storage_uri VirtualMachine#storage_uri}.
	StorageUri *string `field:"required" json:"storageUri" yaml:"storageUri"`
}

