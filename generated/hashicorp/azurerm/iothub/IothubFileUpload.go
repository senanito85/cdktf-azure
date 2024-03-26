package iothub


type IothubFileUpload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#connection_string Iothub#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#container_name Iothub#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#default_ttl Iothub#default_ttl}.
	DefaultTtl *string `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#lock_duration Iothub#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#max_delivery_count Iothub#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"optional" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#notifications Iothub#notifications}.
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iothub#sas_ttl Iothub#sas_ttl}.
	SasTtl *string `field:"optional" json:"sasTtl" yaml:"sasTtl"`
}

