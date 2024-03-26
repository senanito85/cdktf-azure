package eventhub


type EventhubCaptureDescription struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#destination Eventhub#destination}
	Destination *EventhubCaptureDescriptionDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#enabled Eventhub#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#encoding Eventhub#encoding}.
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#interval_in_seconds Eventhub#interval_in_seconds}.
	IntervalInSeconds *float64 `field:"optional" json:"intervalInSeconds" yaml:"intervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#size_limit_in_bytes Eventhub#size_limit_in_bytes}.
	SizeLimitInBytes *float64 `field:"optional" json:"sizeLimitInBytes" yaml:"sizeLimitInBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/eventhub#skip_empty_archives Eventhub#skip_empty_archives}.
	SkipEmptyArchives interface{} `field:"optional" json:"skipEmptyArchives" yaml:"skipEmptyArchives"`
}

