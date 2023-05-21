package jsonrpc

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/edikrisnayana/http-api/router"
	"github.com/gin-gonic/gin"
)

func jsonRPCHandler(ctx *gin.Context, methodMap MapFunc) {
	var request request

	if err := ctx.BindJSON(&request); err != nil {
		newErrorResponse(nil, parseError)
		return
	}

	response := response{}
	if request.IsValid() {
		var jsonErr jsonrpcError
		defer func() {
			if err := recover(); err != nil {
				jsonErr = newError(internalErrorCode, fmt.Sprintf("%v", err), nil)
				response = newErrorResponse(request.Id, jsonErr)
				ctx.IndentedJSON(http.StatusOK, response)
			}
		}()

		method, exist := methodMap[request.Method]
		if exist {
			function := function{
				method: reflect.ValueOf(method),
			}
			args, jsonErr := function.getArgs(request.Params)
			if jsonErr != nil {
				response = newErrorResponse(request.Id, *jsonErr)
			}

			results := function.method.Call(args)
			if len(results) > 0 {
				if len(results) > 1 {
					res := make([]any, len(results))
					for i, result := range results {
						res[i] = result.Interface()
					}
					response = newResponseWithId(request.Id, res)
				} else {
					response = newResponseWithId(request.Id, results[0].Interface())
				}
			} else {
				response = newResponseWithId(request.Id, nil)
			}
		} else {
			response = newErrorResponse(request.Id, methodNotFound)
		}
	} else if request.Id != nil {
		response = newErrorResponse(request.Id, invalidRequest)
	} else {
		response = buildErrorResponse(invalidRequest.Code, "id should not empty", nil)
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

func Register(r *router.Router, endpoint string, methods MapFunc) {
	r.Engine.POST(endpoint, func(ctx *gin.Context) { jsonRPCHandler(ctx, methods) })
}
