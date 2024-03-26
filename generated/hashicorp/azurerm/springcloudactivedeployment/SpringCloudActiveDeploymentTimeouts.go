package springcloudactivedeployment


type SpringCloudActiveDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_active_deployment#create SpringCloudActiveDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_active_deployment#delete SpringCloudActiveDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_active_deployment#read SpringCloudActiveDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_active_deployment#update SpringCloudActiveDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

