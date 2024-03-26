package cosmosdbsqlcontainer


type CosmosdbSqlContainerConflictResolutionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_container#mode CosmosdbSqlContainer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_container#conflict_resolution_path CosmosdbSqlContainer#conflict_resolution_path}.
	ConflictResolutionPath *string `field:"optional" json:"conflictResolutionPath" yaml:"conflictResolutionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_container#conflict_resolution_procedure CosmosdbSqlContainer#conflict_resolution_procedure}.
	ConflictResolutionProcedure *string `field:"optional" json:"conflictResolutionProcedure" yaml:"conflictResolutionProcedure"`
}

