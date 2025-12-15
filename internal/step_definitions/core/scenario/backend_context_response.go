package scenario

func (bc *BackendContext) SetResponse(response *UnifiedResponse) {
	bc.Response = response
}

func (bc *BackendContext) GetResponse() *UnifiedResponse {
	return bc.Response
}

func (bc *BackendContext) HasResponse() bool {
	return bc.Response != nil
}

func (bc *BackendContext) GetResponseBody() []byte {
	if bc.Response == nil {
		return nil
	}
	return bc.Response.Body
}

func (bc *BackendContext) GetStatusCode() int {
	if bc.Response == nil {
		return 0
	}
	return bc.Response.StatusCode
}
