package mediatransform


type MediaTransformOutputAudioAnalyzerPreset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_transform#audio_analysis_mode MediaTransform#audio_analysis_mode}.
	AudioAnalysisMode *string `field:"optional" json:"audioAnalysisMode" yaml:"audioAnalysisMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_transform#audio_language MediaTransform#audio_language}.
	AudioLanguage *string `field:"optional" json:"audioLanguage" yaml:"audioLanguage"`
}

