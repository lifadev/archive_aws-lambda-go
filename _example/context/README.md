# Example - Context

## Execution

```go
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
```

## Interpretation

### Log output

```
START RequestId: df288a66-9ab6-11e6-8f8e-05a606198c73 Version: $LATEST
2016-10-25T13:28:10.012Z	df288a66-9ab6-11e6-8f8e-05a606198c73	Time left before timeout: 2999
2016-10-25T13:28:10.012Z	df288a66-9ab6-11e6-8f8e-05a606198c73	Context:
2016-10-25T13:28:10.012Z	df288a66-9ab6-11e6-8f8e-05a606198c73	{
  "function_name": "bonjour-gopher",
  "function_version": "$LATEST",
  "invoked_function_arn": "arn:aws:lambda:eu-west-1:XXXXXXXXXXXX:function:bonjour-gopher",
  "memory_limit_in_mb": "128",
  "aws_request_id": "df288a66-9ab6-11e6-8f8e-05a606198c73",
  "log_group_name": "/aws/lambda/bonjour-gopher",
  "log_stream_name": "2016/10/25/[$LATEST]5d392c3947074d74ba842a41366c32d9"
}
2016-10-25T13:28:10.512Z	df288a66-9ab6-11e6-8f8e-05a606198c73	Time left before timeout: 2499
END RequestId: df288a66-9ab6-11e6-8f8e-05a606198c73
REPORT RequestId: df288a66-9ab6-11e6-8f8e-05a606198c73	Duration: 500.90 ms	Billed Duration: 600 ms 	Memory Size: 128 MB	Max Memory Used: 8 MB	
```

### Execution result

```json
null
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
