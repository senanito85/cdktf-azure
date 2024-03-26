package webpubsub


type WebPubsubLiveTrace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_pubsub#connectivity_logs_enabled WebPubsub#connectivity_logs_enabled}.
	ConnectivityLogsEnabled interface{} `field:"optional" json:"connectivityLogsEnabled" yaml:"connectivityLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_pubsub#enabled WebPubsub#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_pubsub#http_request_logs_enabled WebPubsub#http_request_logs_enabled}.
	HttpRequestLogsEnabled interface{} `field:"optional" json:"httpRequestLogsEnabled" yaml:"httpRequestLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/web_pubsub#messaging_logs_enabled WebPubsub#messaging_logs_enabled}.
	MessagingLogsEnabled interface{} `field:"optional" json:"messagingLogsEnabled" yaml:"messagingLogsEnabled"`
}

