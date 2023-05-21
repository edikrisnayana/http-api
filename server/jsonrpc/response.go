package jsonrpc

type response struct {
	Jrpc   string        `json:"jsonrpc"`
	Result interface{}   `json:"result,omitempty"`
	Err    *jsonrpcError `json:"error,omitempty"`
	Id     interface{}   `json:"id"`
}

func newResponse(result interface{}) response {
	return newResponseWithId(nil, result)
}

func newResponseWithId(id, result interface{}) response {
	return response{Jrpc: "2.0", Id: id, Result: result}
}

func buildErrorResponse(code int, message string, data interface{}) response {
	return newErrorResponse(nil, newError(code, message, data))
}

func newErrorResponse(id interface{}, err jsonrpcError) response {
	return response{Jrpc: "2.0", Id: id, Err: &err}
}
