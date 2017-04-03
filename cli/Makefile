NAME := vincross/mindcli
ARTIFACTS := \
	bin/mind \
	bin/mind-Darwin-x86_64 \
	bin/mind-Linux-x86_64 \
	bin/mind.exe

INSTALL_BIN := ${GOBIN}/mind

.PHONY : install build push upload clean attach

build : ${ARTIFACTS}

${ARTIFACTS} : main.go
	docker build -t ${NAME} .
	mkdir -p bin
	go build -i -o bin/mind mindsdk/cli 
	GOOS=darwin  GOARCH=amd64 go build -i -o bin/mind-Darwin-x86_64 mindsdk/cli 
	GOOS=linux   GOARCH=amd64 go build -i -o bin/mind-Linux-x86_64 mindsdk/cli 
	GOOS=windows GOARCH=amd64 go build -i -o bin/mind.exe mindsdk/cli

${INSTALL_BIN} : ${ARTIFACTS}
	cp bin/mind ${GOBIN}/mind

install : ${INSTALL_BIN}

uninstall :
	rm -f ${INSTALL_BIN}

push :
	docker push ${NAME}

upload :
	scp ${ARTIFACTS} root@vcrxyz.com:/mindcli

clean : 
	rm -rf bin

attach :
	docker run -it vincross/mindcli bash
