# System build

build:
	go build -buildmode=c-shared -ldflags="-w -s" -o handler.so

pack:
	zip handler.zip handler.so

# Docker build

dbuild:
	@docker run --rm \
		-u ${shell id -u}:${shell id -g} \
		-v ${GOPATH}:/go \
		-w $(subst $(GOPATH),/go,$(PWD)) \
		eawsy/aws-lambda-go make build

dpack:
	@docker run --rm \
		-u ${shell id -u}:${shell id -g} \
		-v ${GOPATH}:/go \
		-w $(subst $(GOPATH),/go,$(PWD)) \
		eawsy/aws-lambda-go make pack

clean:
	@rm -rf handler.zip handler.so
