.PHONY : build build_calculate_lib clean

ARTIFACTS_DIR := $(shell pwd)
ARTIFACTS := ${ARTIFACTS_DIR}/robot/deps/lib/libcalculate.so

build: ${ARTIFACTS}
	mind build

${ARTIFACTS}:
	mind x make build_calculate_lib
	
build_calculate_lib:
	${CC} -fPIC -shared -c ./robot/src/calculate/calculate.c -o ./robot/deps/lib/libcalculate.so

clean:
	rm -f ${ARTIFACTS}
