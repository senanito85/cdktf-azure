package sentinelalertrulescheduled

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SentinelAlertRuleScheduledConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#display_name SentinelAlertRuleScheduled#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#log_analytics_workspace_id SentinelAlertRuleScheduled#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#name SentinelAlertRuleScheduled#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#query SentinelAlertRuleScheduled#query}.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#severity SentinelAlertRuleScheduled#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#alert_rule_template_guid SentinelAlertRuleScheduled#alert_rule_template_guid}.
	AlertRuleTemplateGuid *string `field:"optional" json:"alertRuleTemplateGuid" yaml:"alertRuleTemplateGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#description SentinelAlertRuleScheduled#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#enabled SentinelAlertRuleScheduled#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#event_grouping SentinelAlertRuleScheduled#event_grouping}
	EventGrouping *SentinelAlertRuleScheduledEventGrouping `field:"optional" json:"eventGrouping" yaml:"eventGrouping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#id SentinelAlertRuleScheduled#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incident_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#incident_configuration SentinelAlertRuleScheduled#incident_configuration}
	IncidentConfiguration *SentinelAlertRuleScheduledIncidentConfiguration `field:"optional" json:"incidentConfiguration" yaml:"incidentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#query_frequency SentinelAlertRuleScheduled#query_frequency}.
	QueryFrequency *string `field:"optional" json:"queryFrequency" yaml:"queryFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#query_period SentinelAlertRuleScheduled#query_period}.
	QueryPeriod *string `field:"optional" json:"queryPeriod" yaml:"queryPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#suppression_duration SentinelAlertRuleScheduled#suppression_duration}.
	SuppressionDuration *string `field:"optional" json:"suppressionDuration" yaml:"suppressionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#suppression_enabled SentinelAlertRuleScheduled#suppression_enabled}.
	SuppressionEnabled interface{} `field:"optional" json:"suppressionEnabled" yaml:"suppressionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#tactics SentinelAlertRuleScheduled#tactics}.
	Tactics *[]*string `field:"optional" json:"tactics" yaml:"tactics"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#timeouts SentinelAlertRuleScheduled#timeouts}
	Timeouts *SentinelAlertRuleScheduledTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#trigger_operator SentinelAlertRuleScheduled#trigger_operator}.
	TriggerOperator *string `field:"optional" json:"triggerOperator" yaml:"triggerOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sentinel_alert_rule_scheduled#trigger_threshold SentinelAlertRuleScheduled#trigger_threshold}.
	TriggerThreshold *float64 `field:"optional" json:"triggerThreshold" yaml:"triggerThreshold"`
}

