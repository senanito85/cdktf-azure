package virtualmachine


type VirtualMachineStorageDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#create_option VirtualMachine#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#lun VirtualMachine#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#name VirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#caching VirtualMachine#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#disk_size_gb VirtualMachine#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#managed_disk_id VirtualMachine#managed_disk_id}.
	ManagedDiskId *string `field:"optional" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#managed_disk_type VirtualMachine#managed_disk_type}.
	ManagedDiskType *string `field:"optional" json:"managedDiskType" yaml:"managedDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#vhd_uri VirtualMachine#vhd_uri}.
	VhdUri *string `field:"optional" json:"vhdUri" yaml:"vhdUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/virtual_machine#write_accelerator_enabled VirtualMachine#write_accelerator_enabled}.
	WriteAcceleratorEnabled interface{} `field:"optional" json:"writeAcceleratorEnabled" yaml:"writeAcceleratorEnabled"`
}

