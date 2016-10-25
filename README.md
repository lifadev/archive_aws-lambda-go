[<img src="_asset/powered-by-aws.png" alt="Powered by Amazon Web Services" align="right">][aws-site-url]
[<img src="_asset/created-by-eawsy.png" alt="Created by eawsy" align="right">][eawsy-site-url]

# AWS Lambda - Go
[![Runtime][runtime-badge]][eawsy-runtime-go-url]
[![Api][api-badge]][eawsy-godoc-url]
[![Chat][chat-badge]][eawsy-gitter-url]
![Status][status-badge]
[![License][license-badge]][eawsy-license-url]
<sup>•</sup> <sup>•</sup> <sup>•</sup>
[![Hire us][hire-us-badge]][eawsy-hire-us-url]

[AWS Lambda][aws-lambda-url]<sup>™</sup> lets you run code without provisioning 
or managing servers. This project allows you to run **vanilla Go code** on the 
AWS Lambda platform.


## Usage

### Hello, World!

#### Authoring

```go
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
```

#### Building

> Before continuing, ensure that you have the GCC compiler and header files and 
  libraries for Python development installed on your system. See 
  [Advanced Building](#building-1) section for more details.

```sh
go get -u -d github.com/eawsy/aws-lambda-go/...
go build -buildmode=c-shared -o handler.so
```

#### Packaging

```sh
zip handler.zip handler.so
```

#### Deploying

  1. Create a Lambda function.
  2. Use **`Python 2.7`** as the runtime.
  3. Use **`handler.handle`** as the handler.
  4. Voilà c'est eawsy!

> There are also some [examples][eawsy-examples-url] to play with :tada:.

## Advanced

### Logging

```go
// ...
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  // ...
  log.Println("Hello, World!")
  // ...
}
// ...
```

Logging is as simple as using the **Go [log package][go-log-url]**. You have 
access to all the functions of the log package like 
`Print*`, `Fatal*` or even `Panic*`, and all the logs are available in the Lambda 
function's CloudWatch log group and log stream with the usual format of the AWS 
Lambda logs.

![CloudWatch logs screenshot][eawsy-logs-img]

### Execution

#### Success

```go
// ...
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  // ...
  return "Hello, World!", nil
}
// ...
```

You can **return anything** you want and that can be marshaled by the Go 
[json package][go-json-url].

![Successful function execution screenshot][eawsy-success-img]

#### Failure

```go
// ...
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  // ...
  return nil, errors.New("Oh, Snap!")
}
// ...
```

You can return **any Go error** you want.

![Failed function execution screenshot][eawsy-failure-img]

### Context

You can **access any runtime information** of the 
[runtime.Context][eawsy-godoc-ctx-url] object.  
You can **access any exposed function** of the 
[runtime.Context][eawsy-godoc-ctx-url] object, like 
[`RemainingTimeInMillis`][eawsy-godoc-rtm-url].

```go
// ...
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  // ...
  select {
    case <-time.After(500 * time.Millisecond):
      log.Printf("Remaining time in ms: %d\n", ctx.RemainingTimeInMillis())
  }
  // ...
}
// ...
```

### Building

This project uses [cgo][go-cgo-url] and [Python C extension][python-ext-url] to
provide a seamless integration between AWS Lambda Python 2.7 runtime and Go 
code. This is how we managed to create one and only one binary to deploy on the
AWS Lambda platform.

> **Important**: The output filename **MUST** be `handler.so`.


```sh
# Normal build
go build -buildmode=c-shared -o handler.so

# Size optimized build (~30% gain)
go build -buildmode=c-shared -ldflags="-w -s" -o handler.so
```

Even if it is not visible in the above command line, Go uses the GCC compiler 
and header files and libraries for Python development. Also, you need to have
these dependencies installed on your system.

```sh
# For Debian families
sudo apt-get install build-essentials pkg-config python-dev

# For Redhat families
sudo dnf groupinstall 'Development Tools'
sudo dnf install pkgconfig python-devel
```

For those who have not yet a Linux OS :trollface: or who do not want to install
these dependencies, we also provide a ready to use 
[Docker image][eawsy-docker-url] along with its 
[Dockerfile][eawsy-dockerfile-url]. 

Do not forget to take a tour in the [examples][eawsy-examples-url] to see how it 
works.


## About

[![eawsy][eawsy-logo-img]][eawsy-site-url]

This project is maintained and funded by Alsanium, SAS.

[We][eawsy-site-url] :heart: [AWS][aws-site-url] and open source software. See 
[our other projects][eawsy-github-url], or [hire us][eawsy-hire-us-url] to help 
you build modern applications on AWS.

## License

This product is licensed to you under the Apache License, Version 2.0
(the "License"); you may not use this product except in compliance with the
License. See [LICENSE][eawsy-license-url] and [NOTICE][eawsy-notice-url] for 
more information.

## Trademark

Alsanium, eawsy, the "Created by eawsy" logo, and the "eawsy" logo are 
trademarks of Alsanium, SAS. or its affiliates in France and/or other countries.

Amazon Web Services, the "Powered by Amazon Web Services" logo, and AWS Lambda
are trademarks of Amazon.com, Inc. or its affiliates in the United States and/or 
other countries.

  [eawsy-site-url]: https://eawsy.com
  [eawsy-github-url]: https://github.com/eawsy
  [eawsy-examples-url]: https://github.com/eawsy/aws-lambda-go/tree/master/_example
  [eawsy-dockerfile-url]: https://github.com/eawsy/aws-lambda-go/blob/master/_docker/Dockerfile
  [eawsy-godoc-url]: https://godoc.org/github.com/eawsy/aws-lambda-go/service/lambda/runtime
  [eawsy-godoc-ctx-url]: https://godoc.org/github.com/eawsy/aws-lambda-go/service/lambda/runtime/#Context
  [eawsy-godoc-rtm-url]: https://godoc.org/github.com/eawsy/aws-lambda-go/service/lambda/runtime/#Context.RemainingTimeInMillis
  [eawsy-gitter-url]: https://gitter.im/eawsy/bavardage
  [eawsy-license-url]: https://github.com/eawsy/aws-lambda-go/blob/master/LICENSE
  [eawsy-notice-url]: https://github.com/eawsy/aws-lambda-go/blob/master/NOTICE
  [eawsy-hire-us-url]: https://docs.google.com/forms/d/e/1FAIpQLSfPvn1Dgp95DXfvr3ClPHCNF5abi4D1grveT5btVyBHUk0nXw/viewform
  [eawsy-runtime-go-url]: https://github.com/eawsy/aws-lambda-go
  [eawsy-docker-url]: https://hub.docker.com/r/eawsy/aws-lambda-go/
  [eawsy-logo-img]: _asset/eawsy.png
  [eawsy-logs-img]: _asset/screenshot_logs.png
  [eawsy-success-img]: _asset/screenshot_success.png
  [eawsy-failure-img]: _asset/screenshot_failure.png  
  [go-cgo-url]: https://golang.org/cmd/cgo/
  [go-log-url]: https://golang.org/pkg/log/
  [go-json-url]: https://golang.org/pkg/encoding/json/
  [python-ext-url]: https://docs.python.org/2/extending/extending.html
  [aws-site-url]: https://aws.amazon.com/
  [aws-lambda-url]: https://aws.amazon.com/lambda/  
  [runtime-badge]: http://img.shields.io/badge/runtime-go-ef6c00.svg?style=flat-square
  [api-badge]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
  [chat-badge]: http://img.shields.io/badge/chat-gitter-e91e63.svg?style=flat-square
  [status-badge]: http://img.shields.io/badge/status-stable-689f38.svg?style=flat-square
  [license-badge]: http://img.shields.io/badge/license-apache-757575.svg?style=flat-square
  [hire-us-badge]: http://img.shields.io/badge/hire-eawsy-2196f3.svg?style=flat-square
  
