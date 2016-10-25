package main

import (
	"encoding/json"
	"log"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Hello, %s!", "World")

	var evto interface{}
	json.Unmarshal(evt, &evto)
	evtb, _ := json.MarshalIndent(evto, "", "  ")
	log.Println(string(evtb))

	ctxb, _ := json.MarshalIndent(ctx, "", "  ")
	log.Println(string(ctxb))

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
