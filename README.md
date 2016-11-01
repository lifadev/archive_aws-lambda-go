[<img src="_asset/powered-by-aws.png" alt="Powered by Amazon Web Services" align="right">][aws-home]
[<img src="_asset/created-by-eawsy.png" alt="Created by eawsy" align="right">][eawsy-home]

# eawsy/aws-lambda-go
> A fast and clean way to execute Go on AWS Lambda.

![Runtime][runtime-badge]
[![Api][api-badge]][eawsy-godoc]
[![Chat][chat-badge]][eawsy-gitter]
![Status][status-badge]
[![License][license-badge]](LICENSE)
<sup>•</sup> <sup>•</sup> <sup>•</sup>
[![Hire us][hire-badge]][eawsy-hire-form]

[AWS Lambda][aws-lambda-home] lets you run code without provisioning or managing servers. For now it only supports 
Node.js, Python and Java. This project provides Go support *without spawning a process*.

## Preview

```sh
go get -d github.com/eawsy/aws-lambda-go/...
```

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

```sh
docker run --rm -v $GOPATH:/go -v $PWD:/tmp eawsy/aws-lambda-go
```

```sh
aws lambda create-function \
  --function-name preview-go \
  --runtime python2.7 --handler handler.handle --zip-file fileb://handler.zip \
  --role arn:aws:iam::AWS_ACCOUNT_NUMBER:role/lambda_basic_execution

aws lambda invoke --function-name preview-go output.txt
```

<kbd>![Preview](_asset/eawsy-preview.png)</kbd>

## Documentation

This [wiki](wiki) is the main source of documentation for developers working with or contributing to the 
project.

## About

[![eawsy](_asset/eawsy-logo.png)][eawsy-home]

This project is maintained and funded by Alsanium, SAS.

[We][eawsy-home] :heart: [AWS][aws-home] and open source software. See [our other projects][eawsy-github], or 
[hire us][eawsy-hire-form] to help you build modern applications on AWS.

## License

This product is licensed to you under the Apache License, Version 2.0 (the "License"); you may not use this product 
except in compliance with the License. See [LICENSE](LICENSE) and [NOTICE](NOTICE) for more information.

## Trademark

Alsanium, eawsy, the "Created by eawsy" logo, and the "eawsy" logo are trademarks of Alsanium, SAS. or its affiliates in 
France and/or other countries.

Amazon Web Services, the "Powered by Amazon Web Services" logo, and AWS Lambda are trademarks of Amazon.com, Inc. or its 
affiliates in the United States and/or other countries.

[eawsy-home]: https://eawsy.com
[eawsy-github]: https://github.com/eawsy
[eawsy-gitter]: https://gitter.im/eawsy/bavardage
[eawsy-godoc]: https://godoc.org/github.com/eawsy/aws-lambda-go/service/lambda/runtime
[eawsy-hire-form]: https://docs.google.com/forms/d/e/1FAIpQLSfPvn1Dgp95DXfvr3ClPHCNF5abi4D1grveT5btVyBHUk0nXw/viewform
[aws-home]: https://aws.amazon.com/
[aws-lambda-home]: https://aws.amazon.com/lambda/
[runtime-badge]: http://img.shields.io/badge/runtime-go-ef6c00.svg?style=flat-square
[api-badge]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[chat-badge]: http://img.shields.io/badge/chat-gitter-e91e63.svg?style=flat-square
[status-badge]: http://img.shields.io/badge/status-stable-689f38.svg?style=flat-square
[license-badge]: http://img.shields.io/badge/license-apache-757575.svg?style=flat-square
[hire-badge]: http://img.shields.io/badge/hire-eawsy-2196f3.svg?style=flat-square
