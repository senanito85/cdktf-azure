package servicebustopicauthorizationrule


type ServicebusTopicAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_topic_authorization_rule#create ServicebusTopicAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_topic_authorization_rule#delete ServicebusTopicAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_topic_authorization_rule#read ServicebusTopicAuthorizationRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/servicebus_topic_authorization_rule#update ServicebusTopicAuthorizationRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

