package datafactory


type DataFactoryGithubConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory#account_name DataFactory#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory#branch_name DataFactory#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory#git_url DataFactory#git_url}.
	GitUrl *string `field:"required" json:"gitUrl" yaml:"gitUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory#repository_name DataFactory#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/data_factory#root_folder DataFactory#root_folder}.
	RootFolder *string `field:"required" json:"rootFolder" yaml:"rootFolder"`
}

