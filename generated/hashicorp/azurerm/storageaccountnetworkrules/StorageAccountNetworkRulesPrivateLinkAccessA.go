package storageaccountnetworkrules


type StorageAccountNetworkRulesPrivateLinkAccessA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account_network_rules#endpoint_resource_id StorageAccountNetworkRulesA#endpoint_resource_id}.
	EndpointResourceId *string `field:"required" json:"endpointResourceId" yaml:"endpointResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/storage_account_network_rules#endpoint_tenant_id StorageAccountNetworkRulesA#endpoint_tenant_id}.
	EndpointTenantId *string `field:"optional" json:"endpointTenantId" yaml:"endpointTenantId"`
}

