//go:build no_runtime_type_checking

package customprovider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomProviderActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomProviderActionList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomProviderActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomProviderActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProviderActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProviderActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomProviderActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomProviderActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

