package jsonrpc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type function struct {
	method reflect.Value
}

type MapFunc map[string]any

func (f *function) getArgs(params []any) ([]reflect.Value, *jsonrpcError) {
	args := make([]reflect.Value, len(params))
	for i, param := range params {
		switch param.(type) {
		case float64:
			val, ok := param.(float64)
			if !ok {
				jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v", param), nil)
				return nil, &jsonErr
			}
			arg, jsonErr := convertValue(f.method.Type().In(i), val)
			if jsonErr != nil {
				return nil, jsonErr
			}
			args[i] = arg
		case string:
			var val any
			val, ok := param.(string)
			if !ok {
				jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v", param), nil)
				return nil, &jsonErr
			}

			var jsonErr *jsonrpcError
			if f.method.Type().In(i).Kind() != reflect.String {
				var err error
				if f.method.Type().In(i).Kind() == reflect.Bool {
					val, err = strconv.ParseBool(param.(string))
				} else {
					val, err = strconv.ParseFloat(param.(string), 64)
				}
				if err != nil {
					jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v with error %s", param, err.Error()), nil)
					return nil, &jsonErr
				}
			}

			arg, jsonErr := convertValue(f.method.Type().In(i), val)
			if jsonErr != nil {
				return nil, jsonErr
			}
			args[i] = arg
		case bool:
			val, ok := param.(bool)
			if !ok {
				jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v", param), nil)
				return nil, &jsonErr
			}
			arg, jsonErr := convertValue(f.method.Type().In(i), val)
			if jsonErr != nil {
				return nil, jsonErr
			}
			args[i] = arg
		case map[string]interface{}:
			val := reflect.New(f.method.Type().In(i)).Interface()
			bytes, err := json.Marshal(param)
			if err != nil {
				jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v with error %s", param, err.Error()), nil)
				return nil, &jsonErr
			}
			err = json.Unmarshal(bytes, &val)
			if err != nil {
				jsonErr := newError(invalidRequestCode, fmt.Sprintf("failed convert params %v with error %s", param, err.Error()), nil)
				return nil, &jsonErr
			}
			args[i] = reflect.ValueOf(val).Elem()
		default:
		}

	}

	return args, nil
}
