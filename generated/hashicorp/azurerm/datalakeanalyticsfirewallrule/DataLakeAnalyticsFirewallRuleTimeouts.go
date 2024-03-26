package datalakeanalyticsfirewallrule


type DataLakeAnalyticsFirewallRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_lake_analytics_firewall_rule#create DataLakeAnalyticsFirewallRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_lake_analytics_firewall_rule#delete DataLakeAnalyticsFirewallRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_lake_analytics_firewall_rule#read DataLakeAnalyticsFirewallRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_lake_analytics_firewall_rule#update DataLakeAnalyticsFirewallRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

