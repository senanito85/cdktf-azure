package networkinterfacebackendaddresspoolassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkInterfaceBackendAddressPoolAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_backend_address_pool_association#backend_address_pool_id NetworkInterfaceBackendAddressPoolAssociation#backend_address_pool_id}.
	BackendAddressPoolId *string `field:"required" json:"backendAddressPoolId" yaml:"backendAddressPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_backend_address_pool_association#ip_configuration_name NetworkInterfaceBackendAddressPoolAssociation#ip_configuration_name}.
	IpConfigurationName *string `field:"required" json:"ipConfigurationName" yaml:"ipConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_backend_address_pool_association#network_interface_id NetworkInterfaceBackendAddressPoolAssociation#network_interface_id}.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_backend_address_pool_association#id NetworkInterfaceBackendAddressPoolAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_interface_backend_address_pool_association#timeouts NetworkInterfaceBackendAddressPoolAssociation#timeouts}
	Timeouts *NetworkInterfaceBackendAddressPoolAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

