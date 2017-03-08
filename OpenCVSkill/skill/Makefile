OPENCV_DIR := ../opencv
OPENCV_ARTIFACTS := ${OPENCV_DIR}/artifacts/include ${OPENCV_DIR}/artifacts/lib
DEPS := robot/deps/
SKILL_BIN := robot/skill
MPK := skill.mpk

.PHONY : pack build clean

pack : ${MPK}

${MPK} : ${SKILL_BIN}
	mind pack

build : ${SKILL_BIN}

${SKILL_BIN} : ${DEPS}
	git submodule update && cd robot/src/vendor/github.com/lazywei/go-opencv/ && git init && git checkout 76a80c792afab341300db33893a39c06b366787b
	mind build

${DEPS} : ${OPENCV_ARTIFACTS}
	cp -R ${OPENCV_ARTIFACTS} ${DEPS}

${OPENCV_ARTIFACTS} :
	cd ${OPENCV_DIR} && mind x make

clean :
	rm -f ${SKILL_BIN} ${MPK}
