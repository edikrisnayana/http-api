package main

import (
	"github.com/edikrisnayana/http-api/router"
	"github.com/edikrisnayana/http-api/server/jsonrpc"
)

var address = "localhost:8080"

type request struct {
	A int `json:"a"`
	B int `json:"b"`
}

type response struct {
	Res int `json:"res"`
}

func add(req request) response {
	return response{
		Res: req.A + req.B,
	}
}

func sub(a int, b int) int {
	return a - b
}

func main() {
	router := router.GetRouter(nil)

	jsonrpc.Register(&router, "/math", jsonrpc.MapFunc{
		"add": add,
		"sub": sub,
	})

	router.Start(address)
}
