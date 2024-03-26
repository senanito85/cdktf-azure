package privatednsmxrecord


type PrivateDnsMxRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/private_dns_mx_record#exchange PrivateDnsMxRecord#exchange}.
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/private_dns_mx_record#preference PrivateDnsMxRecord#preference}.
	Preference *float64 `field:"required" json:"preference" yaml:"preference"`
}

