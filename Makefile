GOPATH=$(shell pwd)/gopath/

all: build

run:
	GOPATH=${GOPATH} go run

package: build
	rm -f ${OUTPUTNAME}.7z
	7z a ${OUTPUTNAME}.7z ${OUTPUTNAME}_*

build: get install
	github.com/goreleaser/goreleaser

test:
	GOPATH=${GOPATH} go test .

get:
	GOPATH=${GOPATH} go get golang.org/x/image/font github.com/golang/freetype github.com/goreleaser/goreleaser

release: install
	./gopath/bin/goreleaser

install: get
	GOPATH=${GOPATH} go install github.com/goreleaser/goreleaser