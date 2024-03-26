package sqlserver


type SqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/sql_server#type SqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

