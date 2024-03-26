package cosmosdbaccount


type CosmosdbAccountRestoreDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/cosmosdb_account#collection_names CosmosdbAccount#collection_names}.
	CollectionNames *[]*string `field:"optional" json:"collectionNames" yaml:"collectionNames"`
}

