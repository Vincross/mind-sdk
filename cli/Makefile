NAME := vincross/mindcli
ARTIFACT := ${GOBIN}/mind \
						${GOBIN}/mind-Darwin-x86_64 \
						${GOBIN}/mind-Linux-x86_64 \
						${GOBIN}/mind.exe
MIND_SRC := ../mind

.PHONY : install build clean attach

build : 
	docker build -t ${NAME} .

push :
	docker push ${NAME} 

install : ${ARTIFACT}

${ARTIFACT} : main.go
	GOOS=darwin  GOARCH=amd64  go build -i -o ${GOBIN}/mind-Darwin-x86_64 mindsdk/cli 
	GOOS=darwin  GOARCH=amd64  go build -i -o ${GOBIN}/mind mindsdk/cli 
	GOOS=linux   GOARCH=amd64  go build -i -o ${GOBIN}/mind-Linux-x86_64 mindsdk/cli 
	GOOS=windows GOARCH=amd64  go build -i -o ${GOBIN}/mind.exe mindsdk/cli

upload : ${ARTIFACT}
	scp ${ARTIFACT} root@vcrxyz.com:/mindcli

clean :
	rm -f ${ARTIFACT}

attach :
	docker run -it vincross/mindcli bash
