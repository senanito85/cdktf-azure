//go:build no_runtime_type_checking

package packetcapture

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PacketCaptureFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PacketCaptureFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PacketCaptureFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PacketCaptureFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PacketCaptureFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PacketCaptureFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PacketCaptureFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPacketCaptureFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

