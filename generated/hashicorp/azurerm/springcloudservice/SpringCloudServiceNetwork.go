package springcloudservice


type SpringCloudServiceNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_service#app_subnet_id SpringCloudService#app_subnet_id}.
	AppSubnetId *string `field:"required" json:"appSubnetId" yaml:"appSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_service#cidr_ranges SpringCloudService#cidr_ranges}.
	CidrRanges *[]*string `field:"required" json:"cidrRanges" yaml:"cidrRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_service#service_runtime_subnet_id SpringCloudService#service_runtime_subnet_id}.
	ServiceRuntimeSubnetId *string `field:"required" json:"serviceRuntimeSubnetId" yaml:"serviceRuntimeSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_service#app_network_resource_group SpringCloudService#app_network_resource_group}.
	AppNetworkResourceGroup *string `field:"optional" json:"appNetworkResourceGroup" yaml:"appNetworkResourceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_service#service_runtime_network_resource_group SpringCloudService#service_runtime_network_resource_group}.
	ServiceRuntimeNetworkResourceGroup *string `field:"optional" json:"serviceRuntimeNetworkResourceGroup" yaml:"serviceRuntimeNetworkResourceGroup"`
}

