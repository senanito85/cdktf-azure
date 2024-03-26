//go:build no_runtime_type_checking

package purviewaccount

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PurviewAccountIdentityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PurviewAccountIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PurviewAccountIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PurviewAccountIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPurviewAccountIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

