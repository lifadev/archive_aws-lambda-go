package main

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var s []int

	s[42] = 1337

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
