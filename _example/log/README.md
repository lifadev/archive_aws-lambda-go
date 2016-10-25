# Example - Log

## Execution

```go
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	log.Printf("Hello, %s!", "World")

	return nil, nil
}
```

## Interpretation

### Log output

```
START RequestId: bcba8d4f-9ab1-11e6-b37f-bba61aa6fcfe Version: $LATEST
2016-10-25T12:51:24.88Z	bcba8d4f-9ab1-11e6-b37f-bba61aa6fcfe	Hello, World!
END RequestId: bcba8d4f-9ab1-11e6-b37f-bba61aa6fcfe
REPORT RequestId: bcba8d4f-9ab1-11e6-b37f-bba61aa6fcfe	Duration: 0.77 ms	Billed Duration: 100 ms 	Memory Size: 128 MB	Max Memory Used: 15 MB
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
