package appservicepubliccertificate


type AppServicePublicCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_public_certificate#create AppServicePublicCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_public_certificate#delete AppServicePublicCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_public_certificate#read AppServicePublicCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_public_certificate#update AppServicePublicCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

