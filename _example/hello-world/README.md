# Example - Hello, World!

The simpliest example of running Go code on the AWS Lambda platform.

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
