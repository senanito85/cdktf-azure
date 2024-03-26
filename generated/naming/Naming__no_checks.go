//go:build no_runtime_type_checking

package naming

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Naming) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_Naming) validateAddProviderParameters(provider interface{}) error {
	return nil
}

func (n *jsiiProxy_Naming) validateGetStringParameters(output *string) error {
	return nil
}

func (n *jsiiProxy_Naming) validateInterpolationForOutputParameters(moduleOutput *string) error {
	return nil
}

func (n *jsiiProxy_Naming) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNaming_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNaming_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewNamingParameters(scope constructs.Construct, id *string, config *NamingConfig) error {
	return nil
}

