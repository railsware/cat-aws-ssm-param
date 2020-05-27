build:
	make build-linux
	make build-macos

build-linux: PLATFORM_CONFIG=GOOS=linux GOARCH=amd64
build-linux: OS_NAME=linux
build-linux: build-base

build-macos: PLATFORM_CONFIG=GOOS=darwin GOARCH=amd64
build-macos: OS_NAME=macos
build-macos: build-base

build-base:
	$(PLATFORM_CONFIG) go build -ldflags "-w" -o out/cat-aws-ssm-param.$(OS_NAME)
	upx -9 out/cat-aws-ssm-param.$(OS_NAME)
	gzip out/cat-aws-ssm-param.$(OS_NAME)


clean:
	rm -rf out
