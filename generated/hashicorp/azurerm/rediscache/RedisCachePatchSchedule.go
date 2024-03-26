package rediscache


type RedisCachePatchSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#day_of_week RedisCache#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#maintenance_window RedisCache#maintenance_window}.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#start_hour_utc RedisCache#start_hour_utc}.
	StartHourUtc *float64 `field:"optional" json:"startHourUtc" yaml:"startHourUtc"`
}

