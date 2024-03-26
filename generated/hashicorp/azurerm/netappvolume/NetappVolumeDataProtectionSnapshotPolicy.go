package netappvolume


type NetappVolumeDataProtectionSnapshotPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/netapp_volume#snapshot_policy_id NetappVolume#snapshot_policy_id}.
	SnapshotPolicyId *string `field:"required" json:"snapshotPolicyId" yaml:"snapshotPolicyId"`
}

