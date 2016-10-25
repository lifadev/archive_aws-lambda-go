package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	rt := ctx.RemainingTimeInMillis()
	log.Printf("Time left before timeout: %d", rt)

	log.Println("Context:")
	ctxb, _ := json.MarshalIndent(ctx, "", "  ")
	log.Println(string(ctxb))

	select {
	case <-time.After(500 * time.Millisecond):
		rt = ctx.RemainingTimeInMillis()
		log.Printf("Time left before timeout: %d", rt)
	}

	return nil, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
