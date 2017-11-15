# There are many different ways to create the Makefile to cross compile the lib. We give two of them as follows:
# 1. Run "make build_1", create the libcalculate.so in ./robot/src/calculate/, then move it to ./deps/lib/.
# 2. Run "make build_2", create the libcalculate.so in ./deps/lib directly.
# The point to cross compiling is using the appropriate compiler. We have provided the utils in the mindcli docker image.
# Run "mind x bash" to visit the container, then run "env" you can find the utils.
# Use "mind x make" can help you to run your Makefile with the container of mindcli.

.PHONY : build_1 build_2 calculate_1 calculate_2 clean

all:
	@echo "Try to run \"make build_1\" or \"make build_2\"."

build_1: calculate_1
	mv ./robot/src/calculate/libcalculate.so ./robot/deps/lib
	mind build

calculate_1:
	cd ./robot/src/calculate && mind x make

build_2: calculate_2
	mind build

calculate_2:
	mind x make build_calculate_lib
	
build_calculate_lib:
	${CC} -fPIC -shared -c ./robot/src/calculate/calculate.c -o ./robot/deps/lib/libcalculate.so

clean:
	rm -f ./robot/deps/lib/libcalculate.so
