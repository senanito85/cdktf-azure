package packetcapture

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PacketCaptureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#name PacketCapture#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#network_watcher_name PacketCapture#network_watcher_name}.
	NetworkWatcherName *string `field:"required" json:"networkWatcherName" yaml:"networkWatcherName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#resource_group_name PacketCapture#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// storage_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#storage_location PacketCapture#storage_location}
	StorageLocation *PacketCaptureStorageLocation `field:"required" json:"storageLocation" yaml:"storageLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#target_resource_id PacketCapture#target_resource_id}.
	TargetResourceId *string `field:"required" json:"targetResourceId" yaml:"targetResourceId"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#filter PacketCapture#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#id PacketCapture#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#maximum_bytes_per_packet PacketCapture#maximum_bytes_per_packet}.
	MaximumBytesPerPacket *float64 `field:"optional" json:"maximumBytesPerPacket" yaml:"maximumBytesPerPacket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#maximum_bytes_per_session PacketCapture#maximum_bytes_per_session}.
	MaximumBytesPerSession *float64 `field:"optional" json:"maximumBytesPerSession" yaml:"maximumBytesPerSession"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#maximum_capture_duration PacketCapture#maximum_capture_duration}.
	MaximumCaptureDuration *float64 `field:"optional" json:"maximumCaptureDuration" yaml:"maximumCaptureDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#timeouts PacketCapture#timeouts}
	Timeouts *PacketCaptureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

