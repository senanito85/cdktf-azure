//go:build no_runtime_type_checking

package customprovider

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomProviderResourceTypeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CustomProviderResourceTypeList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CustomProviderResourceTypeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CustomProviderResourceTypeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CustomProviderResourceTypeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CustomProviderResourceTypeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CustomProviderResourceTypeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCustomProviderResourceTypeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

