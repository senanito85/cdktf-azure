package servicebussubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#max_delivery_count ServicebusSubscription#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"required" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#name ServicebusSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#auto_delete_on_idle ServicebusSubscription#auto_delete_on_idle}.
	AutoDeleteOnIdle *string `field:"optional" json:"autoDeleteOnIdle" yaml:"autoDeleteOnIdle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#dead_lettering_on_filter_evaluation_error ServicebusSubscription#dead_lettering_on_filter_evaluation_error}.
	DeadLetteringOnFilterEvaluationError interface{} `field:"optional" json:"deadLetteringOnFilterEvaluationError" yaml:"deadLetteringOnFilterEvaluationError"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#dead_lettering_on_message_expiration ServicebusSubscription#dead_lettering_on_message_expiration}.
	DeadLetteringOnMessageExpiration interface{} `field:"optional" json:"deadLetteringOnMessageExpiration" yaml:"deadLetteringOnMessageExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#default_message_ttl ServicebusSubscription#default_message_ttl}.
	DefaultMessageTtl *string `field:"optional" json:"defaultMessageTtl" yaml:"defaultMessageTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#enable_batched_operations ServicebusSubscription#enable_batched_operations}.
	EnableBatchedOperations interface{} `field:"optional" json:"enableBatchedOperations" yaml:"enableBatchedOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#forward_dead_lettered_messages_to ServicebusSubscription#forward_dead_lettered_messages_to}.
	ForwardDeadLetteredMessagesTo *string `field:"optional" json:"forwardDeadLetteredMessagesTo" yaml:"forwardDeadLetteredMessagesTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#forward_to ServicebusSubscription#forward_to}.
	ForwardTo *string `field:"optional" json:"forwardTo" yaml:"forwardTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#id ServicebusSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#lock_duration ServicebusSubscription#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#namespace_name ServicebusSubscription#namespace_name}.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#requires_session ServicebusSubscription#requires_session}.
	RequiresSession interface{} `field:"optional" json:"requiresSession" yaml:"requiresSession"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#resource_group_name ServicebusSubscription#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#status ServicebusSubscription#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#timeouts ServicebusSubscription#timeouts}
	Timeouts *ServicebusSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#topic_id ServicebusSubscription#topic_id}.
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_subscription#topic_name ServicebusSubscription#topic_name}.
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

