//go:build no_runtime_type_checking

package privateendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateEndpointCustomDnsConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateEndpointCustomDnsConfigsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

