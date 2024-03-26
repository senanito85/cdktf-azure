package trafficmanagerendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TrafficManagerEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#name TrafficManagerEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#profile_name TrafficManagerEndpoint#profile_name}.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#resource_group_name TrafficManagerEndpoint#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#type TrafficManagerEndpoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#custom_header TrafficManagerEndpoint#custom_header}
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#endpoint_location TrafficManagerEndpoint#endpoint_location}.
	EndpointLocation *string `field:"optional" json:"endpointLocation" yaml:"endpointLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#endpoint_status TrafficManagerEndpoint#endpoint_status}.
	EndpointStatus *string `field:"optional" json:"endpointStatus" yaml:"endpointStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#geo_mappings TrafficManagerEndpoint#geo_mappings}.
	GeoMappings *[]*string `field:"optional" json:"geoMappings" yaml:"geoMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#id TrafficManagerEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#min_child_endpoints TrafficManagerEndpoint#min_child_endpoints}.
	MinChildEndpoints *float64 `field:"optional" json:"minChildEndpoints" yaml:"minChildEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#minimum_required_child_endpoints_ipv4 TrafficManagerEndpoint#minimum_required_child_endpoints_ipv4}.
	MinimumRequiredChildEndpointsIpv4 *float64 `field:"optional" json:"minimumRequiredChildEndpointsIpv4" yaml:"minimumRequiredChildEndpointsIpv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#minimum_required_child_endpoints_ipv6 TrafficManagerEndpoint#minimum_required_child_endpoints_ipv6}.
	MinimumRequiredChildEndpointsIpv6 *float64 `field:"optional" json:"minimumRequiredChildEndpointsIpv6" yaml:"minimumRequiredChildEndpointsIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#priority TrafficManagerEndpoint#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// subnet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#subnet TrafficManagerEndpoint#subnet}
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#target TrafficManagerEndpoint#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#target_resource_id TrafficManagerEndpoint#target_resource_id}.
	TargetResourceId *string `field:"optional" json:"targetResourceId" yaml:"targetResourceId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#timeouts TrafficManagerEndpoint#timeouts}
	Timeouts *TrafficManagerEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/traffic_manager_endpoint#weight TrafficManagerEndpoint#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

