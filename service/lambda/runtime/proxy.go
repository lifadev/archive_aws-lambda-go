// +build proxy

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

// #cgo pkg-config: python2
// #cgo CFLAGS: --std=gnu11
// extern long long proxy_get_remaining_time_in_millis();
// extern void proxy_log(char*);
import "C"

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var logger *ilogger

type ilogger struct {
	id string
}

func (l *ilogger) Write(info []byte) (int, error) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	C.proxy_log(C.CString(fmt.Sprintf("%s\t%s\t%s", now, l.id, string(info))))
	return len(info), nil
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
	ctx := &Context{}

	if err := json.Unmarshal([]byte(C.GoString(rctx)), ctx); err != nil {
		return nil, C.CString(err.Error())
	}

	ctx.RemainingTimeInMillis = func() int64 {
		return int64(C.proxy_get_remaining_time_in_millis())
	}

	logger.id = ctx.AWSRequestID

	var env map[string]string
	if err := json.Unmarshal([]byte(C.GoString(renv)), &env); err != nil {
		return nil, C.CString(err.Error())
	}
	for k, v := range env {
		os.Setenv(k, v)
	}

	if res, err := handler.HandleLambda(evt, ctx); err != nil {
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

func init() {
	logger = &ilogger{"00000000-0000-0000-0000-000000000000"}
	log.SetFlags(0)
	log.SetOutput(logger)
}
