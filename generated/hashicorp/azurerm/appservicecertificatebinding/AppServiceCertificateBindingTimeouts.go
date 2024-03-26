package appservicecertificatebinding


type AppServiceCertificateBindingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_certificate_binding#create AppServiceCertificateBinding#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_certificate_binding#delete AppServiceCertificateBinding#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_certificate_binding#read AppServiceCertificateBinding#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

