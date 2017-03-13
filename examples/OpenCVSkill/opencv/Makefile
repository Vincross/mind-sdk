# This makefile should be executed from within the 
# the cross-compiling container. 
#
# Downloads, unpacks, compiles OpenCV for ARM architecture 
# and installs the artifacts to the ./artifacts folder.
#
# Example:
# $ mind x make
# 
NAME := opencv 
OPENCV_VERSION := 2.4.13.2
OPENCV_URL := https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.tar.gz
OPENCV_ARCHIVE := ${OPENCV_VERSION}.tar.gz
OPENCV_SRC := opencv-${OPENCV_VERSION}
OPENCV_CMAKE_FILE := ${OPENCV_SRC}/platforms/linux/arm-gnueabi.toolchain.cmake
ARTIFACTS_DIR := $(shell pwd)/artifacts
ARTIFACTS := ${ARTIFACTS_DIR}/lib/libopencv_core.so.2.4.13
BUILD_DIR := opencv-${OPENCV_VERSION}/platforms/linux/build_hardfp

.PHONY : build clean

build : ${ARTIFACTS}

${ARTIFACTS} : ${BUILD_DIR} ${OPENCV_CMAKE_FILE}
	apt-get update && apt-get install -y libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libjasper-dev libdc1394-22-dev
	cd ${BUILD_DIR} && cmake -DCMAKE_INSTALL_PREFIX=${ARTIFACTS_DIR} -DCMAKE_TOOLCHAIN_FILE=../arm-gnueabi.toolchain.cmake ../../.. && make && make install

${BUILD_DIR} :
	mkdir -p ${BUILD_DIR}

${OPENCV_CMAKE_FILE} : ${OPENCV_ARCHIVE}
	tar -xzvf ${OPENCV_ARCHIVE}
	touch ${OPENCV_CMAKE_FILE}

${OPENCV_ARCHIVE} : 
	wget ${OPENCV_URL}

clean :
	rm -fr ${ARTIFACTS_DIR}
