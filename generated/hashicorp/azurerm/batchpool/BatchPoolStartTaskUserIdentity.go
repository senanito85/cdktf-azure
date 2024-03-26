package batchpool


type BatchPoolStartTaskUserIdentity struct {
	// auto_user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#auto_user BatchPool#auto_user}
	AutoUser *BatchPoolStartTaskUserIdentityAutoUser `field:"optional" json:"autoUser" yaml:"autoUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

