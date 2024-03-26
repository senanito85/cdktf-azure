package cosmosdbsqlstoredprocedure


type CosmosdbSqlStoredProcedureTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_stored_procedure#create CosmosdbSqlStoredProcedure#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_stored_procedure#delete CosmosdbSqlStoredProcedure#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_stored_procedure#read CosmosdbSqlStoredProcedure#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_stored_procedure#update CosmosdbSqlStoredProcedure#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

