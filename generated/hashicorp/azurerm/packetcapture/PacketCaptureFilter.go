package packetcapture


type PacketCaptureFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#protocol PacketCapture#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#local_ip_address PacketCapture#local_ip_address}.
	LocalIpAddress *string `field:"optional" json:"localIpAddress" yaml:"localIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#local_port PacketCapture#local_port}.
	LocalPort *string `field:"optional" json:"localPort" yaml:"localPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#remote_ip_address PacketCapture#remote_ip_address}.
	RemoteIpAddress *string `field:"optional" json:"remoteIpAddress" yaml:"remoteIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/packet_capture#remote_port PacketCapture#remote_port}.
	RemotePort *string `field:"optional" json:"remotePort" yaml:"remotePort"`
}

