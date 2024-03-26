package logicappstandard


type LogicAppStandardConnectionString struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logic_app_standard#name LogicAppStandard#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logic_app_standard#type LogicAppStandard#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/logic_app_standard#value LogicAppStandard#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

