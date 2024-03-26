package mediaassetfilter


type MediaAssetFilterTrackSelection struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/media_asset_filter#condition MediaAssetFilter#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
}

