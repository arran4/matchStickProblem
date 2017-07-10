GOPATH=$(shell pwd)/gopath/
TAG=$(shell expr `git tag | wc -l ` + 1)

all: build

run:
	GOPATH=${GOPATH} go run main.go

package: build
	rm -f ${OUTPUTNAME}.7z
	7z a ${OUTPUTNAME}.7z ${OUTPUTNAME}_*

build: get install
	github.com/goreleaser/goreleaser

test:
	GOPATH=${GOPATH} go test .

get:
	GOPATH=${GOPATH} go get golang.org/x/image/font github.com/golang/freetype github.com/goreleaser/goreleaser

tag:
	git tag -a ${TAG} -m "${TAG} release"

releaser:
	./gopath/bin/goreleaser

snapshotreleaser:
	./gopath/bin/goreleaser --skip-validate --skip-publish

install: get
	GOPATH=${GOPATH} go install github.com/goreleaser/goreleaser