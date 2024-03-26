package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyLifetimeActionTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_certificate#days_before_expiry KeyVaultCertificate#days_before_expiry}.
	DaysBeforeExpiry *float64 `field:"optional" json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/key_vault_certificate#lifetime_percentage KeyVaultCertificate#lifetime_percentage}.
	LifetimePercentage *float64 `field:"optional" json:"lifetimePercentage" yaml:"lifetimePercentage"`
}

