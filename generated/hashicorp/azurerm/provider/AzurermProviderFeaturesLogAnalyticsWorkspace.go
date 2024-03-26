package provider


type AzurermProviderFeaturesLogAnalyticsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs#permanently_delete_on_destroy AzurermProvider#permanently_delete_on_destroy}.
	PermanentlyDeleteOnDestroy interface{} `field:"required" json:"permanentlyDeleteOnDestroy" yaml:"permanentlyDeleteOnDestroy"`
}

