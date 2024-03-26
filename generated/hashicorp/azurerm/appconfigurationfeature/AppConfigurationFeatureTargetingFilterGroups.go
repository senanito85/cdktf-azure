package appconfigurationfeature


type AppConfigurationFeatureTargetingFilterGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_configuration_feature#name AppConfigurationFeature#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_configuration_feature#rollout_percentage AppConfigurationFeature#rollout_percentage}.
	RolloutPercentage *float64 `field:"required" json:"rolloutPercentage" yaml:"rolloutPercentage"`
}

