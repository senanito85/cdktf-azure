//go:build no_runtime_type_checking

package synapseworkspace

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SynapseWorkspaceIdentityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SynapseWorkspaceIdentityList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SynapseWorkspaceIdentityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SynapseWorkspaceIdentityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SynapseWorkspaceIdentityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SynapseWorkspaceIdentityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSynapseWorkspaceIdentityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

