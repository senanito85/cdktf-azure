//go:build no_runtime_type_checking

package iothubdps

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IothubDpsLinkedHubList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IothubDpsLinkedHubList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IothubDpsLinkedHubList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IothubDpsLinkedHubList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IothubDpsLinkedHubList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IothubDpsLinkedHubList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IothubDpsLinkedHubList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIothubDpsLinkedHubListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

