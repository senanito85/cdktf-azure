package managementgroupsubscriptionassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagementGroupSubscriptionAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/management_group_subscription_association#management_group_id ManagementGroupSubscriptionAssociation#management_group_id}.
	ManagementGroupId *string `field:"required" json:"managementGroupId" yaml:"managementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/management_group_subscription_association#subscription_id ManagementGroupSubscriptionAssociation#subscription_id}.
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/management_group_subscription_association#id ManagementGroupSubscriptionAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/management_group_subscription_association#timeouts ManagementGroupSubscriptionAssociation#timeouts}
	Timeouts *ManagementGroupSubscriptionAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

