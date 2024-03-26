//go:build no_runtime_type_checking

package sqlserver

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlServerExtendedAuditingPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlServerExtendedAuditingPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

