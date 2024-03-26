package springcloudcertificate


type SpringCloudCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_certificate#create SpringCloudCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_certificate#delete SpringCloudCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/spring_cloud_certificate#read SpringCloudCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

