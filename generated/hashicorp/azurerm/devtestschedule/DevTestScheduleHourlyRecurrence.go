package devtestschedule


type DevTestScheduleHourlyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/dev_test_schedule#minute DevTestSchedule#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}

