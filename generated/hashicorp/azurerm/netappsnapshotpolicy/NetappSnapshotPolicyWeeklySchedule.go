package netappsnapshotpolicy


type NetappSnapshotPolicyWeeklySchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_snapshot_policy#days_of_week NetappSnapshotPolicy#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_snapshot_policy#hour NetappSnapshotPolicy#hour}.
	Hour *float64 `field:"required" json:"hour" yaml:"hour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_snapshot_policy#minute NetappSnapshotPolicy#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_snapshot_policy#snapshots_to_keep NetappSnapshotPolicy#snapshots_to_keep}.
	SnapshotsToKeep *float64 `field:"required" json:"snapshotsToKeep" yaml:"snapshotsToKeep"`
}

