//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package runtime

import "C"

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	handler Handler
)

// Handler responds to a Lambda function invocation.
//
// HandleLambda can optionally return a value or an error. What happens to the
// returned value depends on the invocation type used when invoking the Lambda
// function. If the Lambda function returns an error, AWS Lambda recognizes the
// failure and sealizes the error information into JSON and returns it. How to
// get the error information back depends to the invocation type that client
// specifies at the time of function invocation.
//
// If HandleLambda panics, the handler (the caller of HandleLambda) recovers the
// panic, logs a stack trace to the CloudWatch log stream, and terminate the
// Lambda function execution.
type Handler interface {
	HandleLambda(json.RawMessage, *Context) (interface{}, error)
}

// HandlerFunc type is an adapter to allow the use of ordinary functions as
// Lambda handlers. If f is a function with the appropriate signature,
// HandlerFunc(f) is a Handler that calls f.
type HandlerFunc func(json.RawMessage, *Context) (interface{}, error)

// HandleLambda calls f(evt, ctx)
func (f HandlerFunc) HandleLambda(evt json.RawMessage, ctx *Context) (interface{}, error) {
	return f(evt, ctx)
}

// Handle registers the given handler.
func Handle(h Handler) {
	handler = h
}

// HandleFunc registers the given handler function.
func HandleFunc(h func(json.RawMessage, *Context) (interface{}, error)) {
	Handle(HandlerFunc(h))
}

//export handle
func handle(revt, rctx, renv *C.char) (rres *C.char, rerr *C.char) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("%s\n%s", err, buf)
			rres = nil
			rerr = C.CString(fmt.Sprintf("%s", err))
		}
	}()

	evt := json.RawMessage([]byte(C.GoString(revt)))

	var ctx Context
	if err := json.Unmarshal([]byte(C.GoString(rctx)), &ctx); err != nil {
		return nil, C.CString(err.Error())
	}

	log.SetFlags(0)
	log.SetOutput(&ctxLogger{&ctx})

	var env map[string]string
	if err := json.Unmarshal([]byte(C.GoString(renv)), &env); err != nil {
		return nil, C.CString(err.Error())
	}
	for k, v := range env {
		os.Setenv(k, v)
	}

	if res, err := handler.HandleLambda(evt, &ctx); err != nil {
		return nil, C.CString(err.Error())
	} else if res != nil {
		tmp, err := json.Marshal(res)
		if err != nil {
			return nil, C.CString(err.Error())
		}
		return C.CString(string(tmp)), nil
	}
	return nil, nil
}
