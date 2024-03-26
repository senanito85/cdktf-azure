package dnsmxrecord


type DnsMxRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/dns_mx_record#exchange DnsMxRecord#exchange}.
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/dns_mx_record#preference DnsMxRecord#preference}.
	Preference *string `field:"required" json:"preference" yaml:"preference"`
}

