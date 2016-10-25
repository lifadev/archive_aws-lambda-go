# Example - Log

Explore event and context with the standard Go log package.

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
