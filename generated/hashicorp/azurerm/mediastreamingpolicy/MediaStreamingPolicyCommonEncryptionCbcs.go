package mediastreamingpolicy


type MediaStreamingPolicyCommonEncryptionCbcs struct {
	// default_content_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_policy#default_content_key MediaStreamingPolicy#default_content_key}
	DefaultContentKey *MediaStreamingPolicyCommonEncryptionCbcsDefaultContentKey `field:"optional" json:"defaultContentKey" yaml:"defaultContentKey"`
	// drm_fairplay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_policy#drm_fairplay MediaStreamingPolicy#drm_fairplay}
	DrmFairplay *MediaStreamingPolicyCommonEncryptionCbcsDrmFairplay `field:"optional" json:"drmFairplay" yaml:"drmFairplay"`
	// enabled_protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_streaming_policy#enabled_protocols MediaStreamingPolicy#enabled_protocols}
	EnabledProtocols *MediaStreamingPolicyCommonEncryptionCbcsEnabledProtocols `field:"optional" json:"enabledProtocols" yaml:"enabledProtocols"`
}

