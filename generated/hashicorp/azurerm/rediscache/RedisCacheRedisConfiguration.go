package rediscache


type RedisCacheRedisConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#aof_backup_enabled RedisCache#aof_backup_enabled}.
	AofBackupEnabled interface{} `field:"optional" json:"aofBackupEnabled" yaml:"aofBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#aof_storage_connection_string_0 RedisCache#aof_storage_connection_string_0}.
	AofStorageConnectionString0 *string `field:"optional" json:"aofStorageConnectionString0" yaml:"aofStorageConnectionString0"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#aof_storage_connection_string_1 RedisCache#aof_storage_connection_string_1}.
	AofStorageConnectionString1 *string `field:"optional" json:"aofStorageConnectionString1" yaml:"aofStorageConnectionString1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#enable_authentication RedisCache#enable_authentication}.
	EnableAuthentication interface{} `field:"optional" json:"enableAuthentication" yaml:"enableAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#maxfragmentationmemory_reserved RedisCache#maxfragmentationmemory_reserved}.
	MaxfragmentationmemoryReserved *float64 `field:"optional" json:"maxfragmentationmemoryReserved" yaml:"maxfragmentationmemoryReserved"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#maxmemory_delta RedisCache#maxmemory_delta}.
	MaxmemoryDelta *float64 `field:"optional" json:"maxmemoryDelta" yaml:"maxmemoryDelta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#maxmemory_policy RedisCache#maxmemory_policy}.
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#maxmemory_reserved RedisCache#maxmemory_reserved}.
	MaxmemoryReserved *float64 `field:"optional" json:"maxmemoryReserved" yaml:"maxmemoryReserved"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#notify_keyspace_events RedisCache#notify_keyspace_events}.
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#rdb_backup_enabled RedisCache#rdb_backup_enabled}.
	RdbBackupEnabled interface{} `field:"optional" json:"rdbBackupEnabled" yaml:"rdbBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#rdb_backup_frequency RedisCache#rdb_backup_frequency}.
	RdbBackupFrequency *float64 `field:"optional" json:"rdbBackupFrequency" yaml:"rdbBackupFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#rdb_backup_max_snapshot_count RedisCache#rdb_backup_max_snapshot_count}.
	RdbBackupMaxSnapshotCount *float64 `field:"optional" json:"rdbBackupMaxSnapshotCount" yaml:"rdbBackupMaxSnapshotCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/redis_cache#rdb_storage_connection_string RedisCache#rdb_storage_connection_string}.
	RdbStorageConnectionString *string `field:"optional" json:"rdbStorageConnectionString" yaml:"rdbStorageConnectionString"`
}

