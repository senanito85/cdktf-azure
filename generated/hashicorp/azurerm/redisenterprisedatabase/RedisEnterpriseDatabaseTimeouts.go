package redisenterprisedatabase


type RedisEnterpriseDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_enterprise_database#create RedisEnterpriseDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_enterprise_database#delete RedisEnterpriseDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_enterprise_database#read RedisEnterpriseDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

