//go:build no_runtime_type_checking

package mediatransform

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MediaTransformOutputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MediaTransformOutputList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MediaTransformOutputList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MediaTransformOutputList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MediaTransformOutputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MediaTransformOutputList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MediaTransformOutputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMediaTransformOutputListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

