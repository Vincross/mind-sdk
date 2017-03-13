NAME := vincross/mindcli

GOBIN ?= ${GOPATH}/bin

ARTIFACTS := \
	${GOBIN}/mind \
	${GOBIN}/mind-Darwin-x86_64 \
	${GOBIN}/mind-Linux-x86_64 \
	${GOBIN}/mind.exe


.PHONY : install build clean attach

build : 
	docker build -t ${NAME} .

push :
	docker push ${NAME} 

install : main.go
	go build -i -o ${GOBIN}/mind mindsdk/cli 
	GOOS=darwin  GOARCH=amd64 go build -i -o ${GOBIN}/mind-Darwin-x86_64 mindsdk/cli 
	GOOS=linux   GOARCH=amd64 go build -i -o ${GOBIN}/mind-Linux-x86_64 mindsdk/cli 
	GOOS=windows GOARCH=amd64 go build -i -o ${GOBIN}/mind.exe mindsdk/cli

upload : ${ARTIFACTS}
	scp ${ARTIFACTS} root@vcrxyz.com:/mindcli

clean :
	rm -f ${ARTIFACTS}

attach :
	docker run -it vincross/mindcli bash
