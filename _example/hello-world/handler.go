package main

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return "Hello, World!", nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
