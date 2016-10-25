# Example - Hello, World!

## Execution

```go
func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	return "Hello, World!", nil
}
```

## Interpretation

### Log output

```
START RequestId: 756135c8-9ab2-11e6-8b56-7329ce4f0466 Version: $LATEST
END RequestId: 756135c8-9ab2-11e6-8b56-7329ce4f0466
REPORT RequestId: 756135c8-9ab2-11e6-8b56-7329ce4f0466	Duration: 0.73 ms	Billed Duration: 100 ms 	Memory Size: 128 MB	Max Memory Used: 9 MB
```

### Execution result

```json
"Hello, World!"
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
