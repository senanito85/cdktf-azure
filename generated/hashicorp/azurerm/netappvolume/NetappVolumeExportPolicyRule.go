package netappvolume


type NetappVolumeExportPolicyRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#allowed_clients NetappVolume#allowed_clients}.
	AllowedClients *[]*string `field:"required" json:"allowedClients" yaml:"allowedClients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#rule_index NetappVolume#rule_index}.
	RuleIndex *float64 `field:"required" json:"ruleIndex" yaml:"ruleIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#cifs_enabled NetappVolume#cifs_enabled}.
	CifsEnabled interface{} `field:"optional" json:"cifsEnabled" yaml:"cifsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#nfsv3_enabled NetappVolume#nfsv3_enabled}.
	Nfsv3Enabled interface{} `field:"optional" json:"nfsv3Enabled" yaml:"nfsv3Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#nfsv4_enabled NetappVolume#nfsv4_enabled}.
	Nfsv4Enabled interface{} `field:"optional" json:"nfsv4Enabled" yaml:"nfsv4Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#protocols_enabled NetappVolume#protocols_enabled}.
	ProtocolsEnabled *[]*string `field:"optional" json:"protocolsEnabled" yaml:"protocolsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#root_access_enabled NetappVolume#root_access_enabled}.
	RootAccessEnabled interface{} `field:"optional" json:"rootAccessEnabled" yaml:"rootAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#unix_read_only NetappVolume#unix_read_only}.
	UnixReadOnly interface{} `field:"optional" json:"unixReadOnly" yaml:"unixReadOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#unix_read_write NetappVolume#unix_read_write}.
	UnixReadWrite interface{} `field:"optional" json:"unixReadWrite" yaml:"unixReadWrite"`
}

