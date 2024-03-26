package backuppolicyvm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPolicyVmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#backup BackupPolicyVm#backup}
	Backup *BackupPolicyVmBackup `field:"required" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#name BackupPolicyVm#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#recovery_vault_name BackupPolicyVm#recovery_vault_name}.
	RecoveryVaultName *string `field:"required" json:"recoveryVaultName" yaml:"recoveryVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#resource_group_name BackupPolicyVm#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#id BackupPolicyVm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#instant_restore_retention_days BackupPolicyVm#instant_restore_retention_days}.
	InstantRestoreRetentionDays *float64 `field:"optional" json:"instantRestoreRetentionDays" yaml:"instantRestoreRetentionDays"`
	// retention_daily block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#retention_daily BackupPolicyVm#retention_daily}
	RetentionDaily *BackupPolicyVmRetentionDaily `field:"optional" json:"retentionDaily" yaml:"retentionDaily"`
	// retention_monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#retention_monthly BackupPolicyVm#retention_monthly}
	RetentionMonthly *BackupPolicyVmRetentionMonthly `field:"optional" json:"retentionMonthly" yaml:"retentionMonthly"`
	// retention_weekly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#retention_weekly BackupPolicyVm#retention_weekly}
	RetentionWeekly *BackupPolicyVmRetentionWeekly `field:"optional" json:"retentionWeekly" yaml:"retentionWeekly"`
	// retention_yearly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#retention_yearly BackupPolicyVm#retention_yearly}
	RetentionYearly *BackupPolicyVmRetentionYearly `field:"optional" json:"retentionYearly" yaml:"retentionYearly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#tags BackupPolicyVm#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#timeouts BackupPolicyVm#timeouts}
	Timeouts *BackupPolicyVmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/backup_policy_vm#timezone BackupPolicyVm#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

