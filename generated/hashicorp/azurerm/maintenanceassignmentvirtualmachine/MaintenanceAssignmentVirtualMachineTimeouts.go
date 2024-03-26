package maintenanceassignmentvirtualmachine


type MaintenanceAssignmentVirtualMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/maintenance_assignment_virtual_machine#create MaintenanceAssignmentVirtualMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/maintenance_assignment_virtual_machine#delete MaintenanceAssignmentVirtualMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/maintenance_assignment_virtual_machine#read MaintenanceAssignmentVirtualMachine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

