//go:build no_runtime_type_checking

package nullmodule

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Nullmodule) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_Nullmodule) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (n *jsiiProxy_Nullmodule) validateGetStringParameters(output *string) error {
	return nil
}

func (n *jsiiProxy_Nullmodule) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (n *jsiiProxy_Nullmodule) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNullmodule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNullmodule_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Nullmodule) validateSetTriggerParameters(val interface{}) error {
	return nil
}

func validateNewNullmoduleParameters(scope constructs.Construct, id *string, config *NullmoduleConfig) error {
	return nil
}

