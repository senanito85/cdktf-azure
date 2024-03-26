package naming

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NamingConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// It is not recommended that you use prefix by azure you should be using a suffix for your resources.
	Prefix *[]*string `field:"optional" json:"prefix" yaml:"prefix"`
	// It is recommended that you specify a suffix for consistency.
	//
	// please use only lowercase characters when possible.
	Suffix *[]*string `field:"optional" json:"suffix" yaml:"suffix"`
	// If you want to include numbers in the unique generation true.
	UniqueIncludeNumbers *bool `field:"optional" json:"uniqueIncludeNumbers" yaml:"uniqueIncludeNumbers"`
	// Max length of the uniqueness suffix to be added 4.
	UniqueLength *float64 `field:"optional" json:"uniqueLength" yaml:"uniqueLength"`
	// Custom value for the random characters to be used.
	UniqueSeed *string `field:"optional" json:"uniqueSeed" yaml:"uniqueSeed"`
}

