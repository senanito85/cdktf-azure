package servicebussubscriptionrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusSubscriptionRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#filter_type ServicebusSubscriptionRule#filter_type}.
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#name ServicebusSubscriptionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#action ServicebusSubscriptionRule#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// correlation_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#correlation_filter ServicebusSubscriptionRule#correlation_filter}
	CorrelationFilter *ServicebusSubscriptionRuleCorrelationFilter `field:"optional" json:"correlationFilter" yaml:"correlationFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#id ServicebusSubscriptionRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#namespace_name ServicebusSubscriptionRule#namespace_name}.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#resource_group_name ServicebusSubscriptionRule#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#sql_filter ServicebusSubscriptionRule#sql_filter}.
	SqlFilter *string `field:"optional" json:"sqlFilter" yaml:"sqlFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#subscription_id ServicebusSubscriptionRule#subscription_id}.
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#subscription_name ServicebusSubscriptionRule#subscription_name}.
	SubscriptionName *string `field:"optional" json:"subscriptionName" yaml:"subscriptionName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#timeouts ServicebusSubscriptionRule#timeouts}
	Timeouts *ServicebusSubscriptionRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription_rule#topic_name ServicebusSubscriptionRule#topic_name}.
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

