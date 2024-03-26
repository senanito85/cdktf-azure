package firewallpolicy


type FirewallPolicyDns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#network_rule_fqdn_enabled FirewallPolicy#network_rule_fqdn_enabled}.
	NetworkRuleFqdnEnabled interface{} `field:"optional" json:"networkRuleFqdnEnabled" yaml:"networkRuleFqdnEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#proxy_enabled FirewallPolicy#proxy_enabled}.
	ProxyEnabled interface{} `field:"optional" json:"proxyEnabled" yaml:"proxyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/firewall_policy#servers FirewallPolicy#servers}.
	Servers *[]*string `field:"optional" json:"servers" yaml:"servers"`
}

