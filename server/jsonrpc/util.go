package jsonrpc

import (
	"fmt"
	"reflect"
)

func convertValue(valType reflect.Type, src any) (reflect.Value, *jsonrpcError) {
	if src == nil {
		return reflect.Zero(valType), nil
	}

	switch valType.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(src), nil
	case reflect.Int:
		return reflect.ValueOf(int(src.(float64))), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(src.(float64))), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(src.(float64))), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(src.(float64))), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(src.(float64))), nil
	case reflect.Uint:
		return reflect.ValueOf(uint(src.(float64))), nil
	case reflect.Uint8:
		return reflect.ValueOf(uint8(src.(float64))), nil
	case reflect.Uint16:
		return reflect.ValueOf(uint16(src.(float64))), nil
	case reflect.Uint32:
		return reflect.ValueOf(uint32(src.(float64))), nil
	case reflect.Uint64:
		return reflect.ValueOf(uint64(src.(float64))), nil
	case reflect.Float32:
		return reflect.ValueOf(float32(src.(float64))), nil
	case reflect.Float64:
		return reflect.ValueOf(src.(float64)), nil
	case reflect.String:
		return reflect.ValueOf(src), nil
	}

	jsonErr := newError(invalidRequestCode, fmt.Sprintf("unsupported type %v", valType.Kind()), nil)
	return reflect.Zero(valType), &jsonErr
}
