NAME := vincross/mindcli
ARTIFACTS := \
	${GOBIN}/mind \
	${GOBIN}/mind-Darwin-x86_64 \
	${GOBIN}/mind-Linux-x86_64 \
	${GOBIN}/mind.exe

.PHONY : install build clean attach

build : main.go
	docker build -t ${NAME} .
	mkdir -p bin
	go build -i -o bin/mind mindsdk/cli 
	GOOS=darwin  GOARCH=amd64 go build -i -o bin/mind-Darwin-x86_64 mindsdk/cli 
	GOOS=linux   GOARCH=amd64 go build -i -o bin/mind-Linux-x86_64 mindsdk/cli 
	GOOS=windows GOARCH=amd64 go build -i -o bin/mind.exe mindsdk/cli

push :
	docker push ${NAME} 

install : bin/mind 
	cp bin/* ${GOBIN}

upload : ${ARTIFACTS}
	scp ${ARTIFACTS} root@vcrxyz.com:/mindcli

clean :
	rm -rf bin

uninstall :
	rm -f ${ARTIFACTS}

attach :
	docker run -it vincross/mindcli bash
