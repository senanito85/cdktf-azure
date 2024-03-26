package springcloudjavadeployment


type SpringCloudJavaDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_java_deployment#cpu SpringCloudJavaDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_java_deployment#memory SpringCloudJavaDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

