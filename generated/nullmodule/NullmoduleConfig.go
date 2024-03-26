package nullmodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NullmoduleConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The trigger value for the `null_resource` resource in this module.
	//
	// one.
	Trigger interface{} `field:"optional" json:"trigger" yaml:"trigger"`
}

