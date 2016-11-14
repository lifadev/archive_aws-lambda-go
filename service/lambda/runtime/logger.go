// +build cgo

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

// extern void proxy_log(char*);
import "C"

import (
	"fmt"
	"log"
	"time"
)

type rawLogger struct{}

func (l *rawLogger) Write(info []byte) (int, error) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	msg := fmt.Sprintf("%s\t%s", now, string(info))
	C.proxy_log(C.CString(msg))
	return len(info), nil
}

type ctxLogger struct{ ctx *Context }

func (l *ctxLogger) Write(info []byte) (int, error) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	msg := fmt.Sprintf("%s\t%s\t%s", now, l.ctx.AWSRequestID, string(info))
	C.proxy_log(C.CString(msg))
	return len(info), nil
}

func init() {
	log.SetFlags(0)
	log.SetOutput(&rawLogger{})
}
