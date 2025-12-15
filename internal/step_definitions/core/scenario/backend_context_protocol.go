package scenario

func (bc *BackendContext) SetProtocol(p APIProtocol) {
	bc.Protocol = p
}

func (bc *BackendContext) GetProtocol() APIProtocol {
	return bc.Protocol
}

func (bc *BackendContext) IsGraphQL() bool {
	return bc.Protocol != nil && bc.Protocol.GetProtocolName() == "GraphQL"
}
func (bc *BackendContext) IsREST() bool {
	return bc.Protocol != nil && bc.Protocol.GetProtocolName() == "REST"
}
