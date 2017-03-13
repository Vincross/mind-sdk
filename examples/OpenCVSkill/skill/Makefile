OPENCV_DIR := ../opencv
OPENCV_ARTIFACTS := ${OPENCV_DIR}/artifacts/include ${OPENCV_DIR}/artifacts/lib
DEPS_DIR := robot/deps
DEPS := robot/deps/lib robot/deps/include
SKILL_BIN := robot/skill
MPK := skill.mpk

.PHONY : pack build clean run

pack : ${MPK}

${MPK} : ${SKILL_BIN}
	mind pack

run : ${MPK}
	mind run

build : ${SKILL_BIN}

${SKILL_BIN} : ${DEPS}
	git submodule update --init --recursive
	cd robot/src/vendor/github.com/lazywei/go-opencv/ && git checkout 76a80c792afab341300db33893a39c06b366787b
	mind build

${DEPS} : ${DEPS_DIR} ${OPENCV_ARTIFACTS}
	cp -R ${OPENCV_ARTIFACTS} ${DEPS_DIR}

${DEPS_DIR} :
	mkdir -p ${DEPS_DIR}

${OPENCV_ARTIFACTS} :
	cd ${OPENCV_DIR} && mind x make

clean :
	rm -f ${SKILL_BIN} ${MPK}
