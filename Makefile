GOPATH=$(shell pwd)/gopath/
GOOSARCHPKG=$(shell pwd)/goosarchpkg/${GOOS}_${GOARCH}
OUTPUTNAME=matchProblem

all: build

run:
	GOPATH=${GOPATH} go run

package: build
	rm -f ${OUTPUTNAME}.7z
	7z a ${OUTPUTNAME}.7z ${OUTPUTNAME}_*

buildwin32: getwin32
	GOOS=windows GOARCH=386 EXT=.exe GOPATH=${GOPATH} make obuild

buildwin64: getwin64
	GOOS=windows GOARCH=amd64 EXT=.exe GOPATH=${GOPATH} make obuild

buildlinux32: getlinux32
	GOOS=linux GOARCH=386 GOPATH=${GOPATH} make obuild

buildlinux64: getlinux64
	GOOS=linux GOARCH=amd64 GOPATH=${GOPATH} make obuild

buildmac32: getmac32
	GOOS=darwin GOARCH=386 GOPATH=${GOPATH} make obuild

buildmac64: getmac64
	GOOS=darwin GOARCH=amd64 GOPATH=${GOPATH} make obuild

build: buildlinux32 buildlinux64 buildwin32 buildwin64 buildmac32 buildmac64

obuild:
	mkdir -p ${GOOSARCHPKG}
	go build -pkgdir ${GOOSARCHPKG} -o ${OUTPUTNAME}_${GOOS}_${GOARCH}${EXT} .

test:
	GOPATH=${GOPATH} go test .

get:
	#GOPATH=${GOPATH} go get -pkgdir ${GOOSARCHPKG} -u

getlinux32:
	GOOS=linux GOARCH=386 make get

getlinux64:
	GOOS=linux GOARCH=amd64 make get

getwin32:
	GOOS=windows GOARCH=386 make get

getwin64:
	GOOS=windows GOARCH=amd64 make get

getmac32:
	GOOS=darwin GOARCH=386 make get

getmac64:
	GOOS=darwin GOARCH=amd64 make get
