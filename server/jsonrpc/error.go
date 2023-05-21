package jsonrpc

type jsonrpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Official JSON-RPC 2.0 Spec Error Codes and Messages
const (
	lowestReservedErrorCode  = -32768
	parseErrorCode           = -32700
	invalidRequestCode       = -32600
	methodNotFoundCode       = -32601
	invalidParamsCode        = -32602
	internalErrorCode        = -32603
	highestReservedErrorCode = -32000

	parseErrorMessage     = "Parse error"
	invalidRequestMessage = "Invalid Request"
	methodNotFoundMessage = "Method not found"
	invalidParamsMessage  = "Invalid params"
	internalErrorMessage  = "Internal error"
)

// Official Errors
var (
	// ParseError is returned to the client if a JSON is not well formed.
	parseError = newError(parseErrorCode, parseErrorMessage, nil)
	// InvalidRequest is returned to the client if a request does not
	// conform to JSON-RPC 2.0 spec
	invalidRequest = newError(invalidRequestCode, invalidRequestMessage, nil)
	// MethodNotFound is returned to the client if a method is called that
	// has not been registered with RegisterMethod()
	methodNotFound = newError(methodNotFoundCode, methodNotFoundMessage, nil)
	// InvalidParams is returned to the client if a method is called with
	// an invalid "params" object. A method's function is responsible for
	// detecting and returning this error.
	invalidParams = newError(invalidParamsCode, invalidParamsMessage, nil)
	// InternalError is returned to the client if a method function returns
	// an invalid response object.
	internalError = newError(internalErrorCode, internalErrorMessage, nil)
)

func newError(code int, message string, data interface{}) jsonrpcError {
	return jsonrpcError{Code: code, Message: message, Data: data}
}
