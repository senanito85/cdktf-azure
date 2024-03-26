package sentinelalertrulemachinelearningbehavioranalytics

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleMachineLearningBehaviorAnalyticsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#alert_rule_template_guid SentinelAlertRuleMachineLearningBehaviorAnalytics#alert_rule_template_guid}.
	AlertRuleTemplateGuid *string `field:"required" json:"alertRuleTemplateGuid" yaml:"alertRuleTemplateGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#log_analytics_workspace_id SentinelAlertRuleMachineLearningBehaviorAnalytics#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#name SentinelAlertRuleMachineLearningBehaviorAnalytics#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#enabled SentinelAlertRuleMachineLearningBehaviorAnalytics#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#id SentinelAlertRuleMachineLearningBehaviorAnalytics#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_machine_learning_behavior_analytics#timeouts SentinelAlertRuleMachineLearningBehaviorAnalytics#timeouts}
	Timeouts *SentinelAlertRuleMachineLearningBehaviorAnalyticsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

