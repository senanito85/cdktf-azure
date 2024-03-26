package networkprofile


type NetworkProfileContainerNetworkInterface struct {
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_profile#ip_configuration NetworkProfile#ip_configuration}
	IpConfiguration interface{} `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/network_profile#name NetworkProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

