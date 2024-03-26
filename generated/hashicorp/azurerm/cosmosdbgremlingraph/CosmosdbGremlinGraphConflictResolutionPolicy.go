package cosmosdbgremlingraph


type CosmosdbGremlinGraphConflictResolutionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_gremlin_graph#mode CosmosdbGremlinGraph#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_gremlin_graph#conflict_resolution_path CosmosdbGremlinGraph#conflict_resolution_path}.
	ConflictResolutionPath *string `field:"optional" json:"conflictResolutionPath" yaml:"conflictResolutionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_gremlin_graph#conflict_resolution_procedure CosmosdbGremlinGraph#conflict_resolution_procedure}.
	ConflictResolutionProcedure *string `field:"optional" json:"conflictResolutionProcedure" yaml:"conflictResolutionProcedure"`
}

