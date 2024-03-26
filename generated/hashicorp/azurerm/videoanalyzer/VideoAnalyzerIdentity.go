package videoanalyzer


type VideoAnalyzerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/video_analyzer#identity_ids VideoAnalyzer#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/video_analyzer#type VideoAnalyzer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

