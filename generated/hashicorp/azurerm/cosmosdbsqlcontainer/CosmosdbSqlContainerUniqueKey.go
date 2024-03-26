package cosmosdbsqlcontainer


type CosmosdbSqlContainerUniqueKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_sql_container#paths CosmosdbSqlContainer#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
}

