package appservice


type AppServiceSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service#allowed_origins AppService#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service#support_credentials AppService#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

