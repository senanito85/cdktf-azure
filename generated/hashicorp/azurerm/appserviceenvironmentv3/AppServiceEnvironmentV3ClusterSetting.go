package appserviceenvironmentv3


type AppServiceEnvironmentV3ClusterSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_environment_v3#name AppServiceEnvironmentV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_environment_v3#value AppServiceEnvironmentV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

