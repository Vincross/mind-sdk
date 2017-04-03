REGISTRY := vincross
NAME := xcompile
DOCKERFILE := Dockerfile
CACHEFILE := .xcompile

.PHONY : build push attach clean

${CACHEFILE} : ${DOCKERFILE}
	docker build -t ${REGISTRY}/${NAME} .
	touch ${CACHEFILE}

build : ${CACHEFILE}

push : ${CACHEFILE}
	docker push ${REGISTRY}/${NAME}

attach : ${CACHEFILE}
	docker run -it ${REGISTRY}/${NAME} bash

clean : 
	rm -f ${CACHEFILE}
