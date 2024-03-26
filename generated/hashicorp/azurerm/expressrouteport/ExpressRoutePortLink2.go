package expressrouteport


type ExpressRoutePortLink2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_port#admin_enabled ExpressRoutePort#admin_enabled}.
	AdminEnabled interface{} `field:"optional" json:"adminEnabled" yaml:"adminEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_port#macsec_cak_keyvault_secret_id ExpressRoutePort#macsec_cak_keyvault_secret_id}.
	MacsecCakKeyvaultSecretId *string `field:"optional" json:"macsecCakKeyvaultSecretId" yaml:"macsecCakKeyvaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_port#macsec_cipher ExpressRoutePort#macsec_cipher}.
	MacsecCipher *string `field:"optional" json:"macsecCipher" yaml:"macsecCipher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/express_route_port#macsec_ckn_keyvault_secret_id ExpressRoutePort#macsec_ckn_keyvault_secret_id}.
	MacsecCknKeyvaultSecretId *string `field:"optional" json:"macsecCknKeyvaultSecretId" yaml:"macsecCknKeyvaultSecretId"`
}

