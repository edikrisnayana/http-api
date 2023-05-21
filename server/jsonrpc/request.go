package jsonrpc

type request struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params,omitempty"`
	Id      any    `json:"id,omitempty"`
}

func (req request) IsValid() bool {
	if req.Id != nil {
		switch req.Id.(type) {
		case float64:
		case string:
		default:
			return false
		}
	} else {
		return false
	}

	return req.Jsonrpc == "2.0" && len(req.Method) > 0
}
