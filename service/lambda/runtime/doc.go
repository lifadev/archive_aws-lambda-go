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

/*
Package runtime allows running standard Go code on the AWS Lambda platform
https://aws.amazon.com/lambda/.

Creating:

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

Building:

	go build -buildmode=c-shared -ldflags "-s -w" -o handler.so

Packaging:

	zip handler.zip handler.so

Deploying:

	Runtime: Python 2.7
	Handler: handler.handle

Voilà!

Take a tour at https://github.com/eawsy/aws-lambda-go
*/
package runtime
