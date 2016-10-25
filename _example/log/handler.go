package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Hello, %s!", "World")

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
