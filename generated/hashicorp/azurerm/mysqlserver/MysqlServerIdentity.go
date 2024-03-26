package mysqlserver


type MysqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mysql_server#type MysqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

