# Example - Panic

## Execution

```go
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var s []int

	s[42] = 1337

	return nil, nil
}
```

## Interpretation

An out of range index is accessed, and Lambda begins *panicking*:
  1. Normal execution of Lambda stops.
	
  2. Stack trace is logged in CloudWatch.
	
    ```
    ...
    go/src/github.com/eawsy/aws-lambda-go/_example/panic/handler.go:12 +0x16
    ...
    ```
		
  3. Error message is returned to the client.
	
    ```
    ...
    "errorMessage": "runtime error: index out of range"
    ...
    ```

### Log output

```
START RequestId: 5d430164-9ab1-11e6-a979-c545bfec6e13 Version: $LATEST
2016-10-25T12:48:45.06Z	5d430164-9ab1-11e6-a979-c545bfec6e13	runtime error: index out of range
goroutine 17 [running, locked to thread]:
github.com/eawsy/aws-lambda-go/service/lambda/runtime.handle.func1(0xc420045e30, 0xc420045e38)
	go/src/github.com/eawsy/aws-lambda-go/service/lambda/runtime/handler.go:75 +0xbe
panic(0x7f3f275fbb60, 0xc420014080)
	go/src/runtime/panic.go:458 +0x247
main.handle(0xc42007e240, 0x236, 0x240, 0xc420080000, 0x0, 0xc4200103f0, 0xc42008c000, 0x7f3f27675810)
	go/src/github.com/eawsy/aws-lambda-go/_example/panic/handler.go:12 +0x16
github.com/eawsy/aws-lambda-go/service/lambda/runtime.HandlerFunc.HandleLambda(0x7f3f27615bf0, 0xc42007e240, 0x236, 0x240, 0xc420080000, 0x0, 0x0, 0x0, 0x0)
	go/src/github.com/eawsy/aws-lambda-go/service/lambda/runtime/handler.go:56 +0x50
github.com/eawsy/aws-lambda-go/service/lambda/runtime.handle(0x12c51e4, 0x7f3f2846c4ec, 0x1244c44, 0x0, 0x0)
	go/src/github.com/eawsy/aws-lambda-go/service/lambda/runtime/handler.go:100 +0x37e
github.com/eawsy/aws-lambda-go/service/lambda/runtime._cgoexpwrap_3f5915b23838_handle(0x12c51e4, 0x7f3f2846c4ec, 0x1244c44, 0x0, 0x0)
	github.com/eawsy/aws-lambda-go/service/lambda/runtime/_obj/_cgo_gotypes.go:94 +0x87
runtime error: index out of range: error
Traceback (most recent call last):
  File "/var/runtime/awslambda/bootstrap.py", line 204, in handle_event_request
    result = request_handler(json_input, context)
error: runtime error: index out of range

END RequestId: 5d430164-9ab1-11e6-a979-c545bfec6e13
REPORT RequestId: 5d430164-9ab1-11e6-a979-c545bfec6e13	Duration: 0.89 ms	Billed Duration: 100 ms 	Memory Size: 128 MB	Max Memory Used: 13 MB
```

### Execution result

```json
{
  "stackTrace": [
    [
      "/var/runtime/awslambda/bootstrap.py",
      204,
      "handle_event_request",
      "result = request_handler(json_input, context)"
    ]
  ],
  "errorType": "error",
  "errorMessage": "runtime error: index out of range"
}
```

## Usage

### Build & Package 

#### With Docker (only option for OSX users)

```sh
make dbuild
# output: handler.so

make dpack
# output: handler.zip
```

#### Without Docker

```sh
make build
# output: handler.so

make pack
# output: handler.zip
```

### Deploy

![Deploy your Lambda function on AWS][eawsy-config-img]

  [eawsy-config-img]: ../../_asset/screenshot_config.png 
