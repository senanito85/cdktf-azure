//go:build no_runtime_type_checking

package networkinterface

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkInterfaceIpConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkInterfaceIpConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkInterfaceIpConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceIpConfigurationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceIpConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceIpConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceIpConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkInterfaceIpConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

