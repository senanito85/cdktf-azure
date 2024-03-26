package batchpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#node_agent_sku_id BatchPool#node_agent_sku_id}.
	NodeAgentSkuId *string `field:"required" json:"nodeAgentSkuId" yaml:"nodeAgentSkuId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#resource_group_name BatchPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// storage_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#storage_image_reference BatchPool#storage_image_reference}
	StorageImageReference *BatchPoolStorageImageReference `field:"required" json:"storageImageReference" yaml:"storageImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#vm_size BatchPool#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#auto_scale BatchPool#auto_scale}
	AutoScale *BatchPoolAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#certificate BatchPool#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#container_configuration BatchPool#container_configuration}
	ContainerConfiguration *BatchPoolContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#display_name BatchPool#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#fixed_scale BatchPool#fixed_scale}
	FixedScale *BatchPoolFixedScale `field:"optional" json:"fixedScale" yaml:"fixedScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#id BatchPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#identity BatchPool#identity}
	Identity *BatchPoolIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#max_tasks_per_node BatchPool#max_tasks_per_node}.
	MaxTasksPerNode *float64 `field:"optional" json:"maxTasksPerNode" yaml:"maxTasksPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#metadata BatchPool#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#network_configuration BatchPool#network_configuration}
	NetworkConfiguration *BatchPoolNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// start_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#start_task BatchPool#start_task}
	StartTask *BatchPoolStartTask `field:"optional" json:"startTask" yaml:"startTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#stop_pending_resize_operation BatchPool#stop_pending_resize_operation}.
	StopPendingResizeOperation interface{} `field:"optional" json:"stopPendingResizeOperation" yaml:"stopPendingResizeOperation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/batch_pool#timeouts BatchPool#timeouts}
	Timeouts *BatchPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

