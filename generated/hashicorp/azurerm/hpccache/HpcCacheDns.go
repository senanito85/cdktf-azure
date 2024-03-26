package hpccache


type HpcCacheDns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hpc_cache#servers HpcCache#servers}.
	Servers *[]*string `field:"required" json:"servers" yaml:"servers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/hpc_cache#search_domain HpcCache#search_domain}.
	SearchDomain *string `field:"optional" json:"searchDomain" yaml:"searchDomain"`
}

