package firewallpolicy


type FirewallPolicyIntrusionDetection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#mode FirewallPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// signature_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#signature_overrides FirewallPolicy#signature_overrides}
	SignatureOverrides interface{} `field:"optional" json:"signatureOverrides" yaml:"signatureOverrides"`
	// traffic_bypass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#traffic_bypass FirewallPolicy#traffic_bypass}
	TrafficBypass interface{} `field:"optional" json:"trafficBypass" yaml:"trafficBypass"`
}

