package diskpooliscsitarget


type DiskPoolIscsiTargetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/disk_pool_iscsi_target#create DiskPoolIscsiTarget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/disk_pool_iscsi_target#delete DiskPoolIscsiTarget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/disk_pool_iscsi_target#read DiskPoolIscsiTarget#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

