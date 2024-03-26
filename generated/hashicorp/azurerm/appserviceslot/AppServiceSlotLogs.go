package appserviceslot


type AppServiceSlotLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#application_logs AppServiceSlot#application_logs}
	ApplicationLogs *AppServiceSlotLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#detailed_error_messages_enabled AppServiceSlot#detailed_error_messages_enabled}.
	DetailedErrorMessagesEnabled interface{} `field:"optional" json:"detailedErrorMessagesEnabled" yaml:"detailedErrorMessagesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#failed_request_tracing_enabled AppServiceSlot#failed_request_tracing_enabled}.
	FailedRequestTracingEnabled interface{} `field:"optional" json:"failedRequestTracingEnabled" yaml:"failedRequestTracingEnabled"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/app_service_slot#http_logs AppServiceSlot#http_logs}
	HttpLogs *AppServiceSlotLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

