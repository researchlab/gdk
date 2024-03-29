.PHONY: list vet fmt default clean
all: list vet fmt default clean 
# ###### 只需要设置下面两个参数, 其它都不需要修改
BINARY="binary"
VERSION=0.0.1
# #####
BUILD=`date +%F`
SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
versionDir="github.com/researchlab/gdk"
gitTag=$(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi) # format:'%H' 完整的commit id
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD) # 去掉--short 就是完整的commit-id
gitTreeState=$(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-s -w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState} -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch}"


PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@echo "build the ${BINARY}"
	@GOOS=linux GOARCH=amd64 go build -ldflags ${ldflags} -o  build/${BINARY}.linux  -tags=jsoniter
	@GOOS=darwin GOARCH=arm64 go build -ldflags ${ldflags} -o  build/${BINARY}.mac  -tags=jsoniter
	@echo "build done."

list:
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@echo "fmt the project"
	@gofmt -s -w ${GOFILES}

vet:
	@echo "check the project codes."
	@go vet $(VETPACKAGES)
	@echo "check done."


clean:
	@rm -rf build/*

