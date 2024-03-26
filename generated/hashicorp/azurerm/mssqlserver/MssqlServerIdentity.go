package mssqlserver


type MssqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server#type MssqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/mssql_server#user_assigned_identity_ids MssqlServer#user_assigned_identity_ids}.
	UserAssignedIdentityIds *[]*string `field:"optional" json:"userAssignedIdentityIds" yaml:"userAssignedIdentityIds"`
}

