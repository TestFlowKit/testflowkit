package scenario

func (bc *BackendContext) GetHeader(name string) (string, bool) {
	value, exists := bc.Headers[name]
	return value, exists
}

func (bc *BackendContext) SetHeader(name, value string) {
	bc.Headers[name] = value
}

func (bc *BackendContext) GetHeaders() map[string]string {
	return bc.Headers
}

func (bc *BackendContext) ClearHeaders() {
	bc.Headers = make(map[string]string)
}
