package devtestschedule


type DevTestScheduleDailyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/dev_test_schedule#time DevTestSchedule#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

