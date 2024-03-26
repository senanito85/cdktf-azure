package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_content_key_policy#playback_duration_seconds MediaContentKeyPolicy#playback_duration_seconds}.
	PlaybackDurationSeconds *float64 `field:"optional" json:"playbackDurationSeconds" yaml:"playbackDurationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_content_key_policy#storage_duration_seconds MediaContentKeyPolicy#storage_duration_seconds}.
	StorageDurationSeconds *float64 `field:"optional" json:"storageDurationSeconds" yaml:"storageDurationSeconds"`
}

