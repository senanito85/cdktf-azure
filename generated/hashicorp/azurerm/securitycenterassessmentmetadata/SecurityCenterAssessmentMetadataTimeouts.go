package securitycenterassessmentmetadata


type SecurityCenterAssessmentMetadataTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/security_center_assessment_metadata#create SecurityCenterAssessmentMetadata#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/security_center_assessment_metadata#delete SecurityCenterAssessmentMetadata#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/security_center_assessment_metadata#read SecurityCenterAssessmentMetadata#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/security_center_assessment_metadata#update SecurityCenterAssessmentMetadata#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

