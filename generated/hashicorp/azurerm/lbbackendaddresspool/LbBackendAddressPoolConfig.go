package lbbackendaddresspool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbBackendAddressPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#loadbalancer_id LbBackendAddressPool#loadbalancer_id}.
	LoadbalancerId *string `field:"required" json:"loadbalancerId" yaml:"loadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#name LbBackendAddressPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// backend_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#backend_address LbBackendAddressPool#backend_address}
	BackendAddress interface{} `field:"optional" json:"backendAddress" yaml:"backendAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#id LbBackendAddressPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#resource_group_name LbBackendAddressPool#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#timeouts LbBackendAddressPool#timeouts}
	Timeouts *LbBackendAddressPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// tunnel_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/lb_backend_address_pool#tunnel_interface LbBackendAddressPool#tunnel_interface}
	TunnelInterface interface{} `field:"optional" json:"tunnelInterface" yaml:"tunnelInterface"`
}

