//go:build no_runtime_type_checking

package kubernetescluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterKubeAdminConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesClusterKubeAdminConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

