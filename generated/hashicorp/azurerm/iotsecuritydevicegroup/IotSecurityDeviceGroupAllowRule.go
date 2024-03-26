package iotsecuritydevicegroup


type IotSecurityDeviceGroupAllowRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#connection_from_ips_not_allowed IotSecurityDeviceGroup#connection_from_ips_not_allowed}.
	ConnectionFromIpsNotAllowed *[]*string `field:"optional" json:"connectionFromIpsNotAllowed" yaml:"connectionFromIpsNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#connection_to_ip_not_allowed IotSecurityDeviceGroup#connection_to_ip_not_allowed}.
	ConnectionToIpNotAllowed *[]*string `field:"optional" json:"connectionToIpNotAllowed" yaml:"connectionToIpNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#connection_to_ips_not_allowed IotSecurityDeviceGroup#connection_to_ips_not_allowed}.
	ConnectionToIpsNotAllowed *[]*string `field:"optional" json:"connectionToIpsNotAllowed" yaml:"connectionToIpsNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#local_user_not_allowed IotSecurityDeviceGroup#local_user_not_allowed}.
	LocalUserNotAllowed *[]*string `field:"optional" json:"localUserNotAllowed" yaml:"localUserNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#local_users_not_allowed IotSecurityDeviceGroup#local_users_not_allowed}.
	LocalUsersNotAllowed *[]*string `field:"optional" json:"localUsersNotAllowed" yaml:"localUsersNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#processes_not_allowed IotSecurityDeviceGroup#processes_not_allowed}.
	ProcessesNotAllowed *[]*string `field:"optional" json:"processesNotAllowed" yaml:"processesNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/2.99.0/docs/resources/iot_security_device_group#process_not_allowed IotSecurityDeviceGroup#process_not_allowed}.
	ProcessNotAllowed *[]*string `field:"optional" json:"processNotAllowed" yaml:"processNotAllowed"`
}

