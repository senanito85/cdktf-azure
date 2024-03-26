package backuppolicyfileshare


type BackupPolicyFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_file_share#create BackupPolicyFileShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_file_share#delete BackupPolicyFileShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_file_share#read BackupPolicyFileShare#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_file_share#update BackupPolicyFileShare#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

