package springcloudapp


type SpringCloudAppPersistentDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_app#size_in_gb SpringCloudApp#size_in_gb}.
	SizeInGb *float64 `field:"required" json:"sizeInGb" yaml:"sizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_app#mount_path SpringCloudApp#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

