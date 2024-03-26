package functionappslot


type FunctionAppSlotAuthSettingsMicrosoft struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/function_app_slot#client_id FunctionAppSlot#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/function_app_slot#client_secret FunctionAppSlot#client_secret}.
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/function_app_slot#oauth_scopes FunctionAppSlot#oauth_scopes}.
	OauthScopes *[]*string `field:"optional" json:"oauthScopes" yaml:"oauthScopes"`
}

