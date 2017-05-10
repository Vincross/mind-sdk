FROM ubuntu:14.04
# Basics and cross-compile tools
ENV CROSS arm-linux-gnueabihf
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    unzip \
    wget \
    git \
    gcc-${CROSS} \
    g++-${CROSS} \
    cmake \
    pkg-config \
    jq \
		ruby-dev \
    && apt-get clean && apt-get autoremove --purge
RUN gem install fpm
# Setup cross compilers
ENV AS=/usr/bin/${CROSS}-as \
    AR=/usr/bin/${CROSS}-ar \
    CC=/usr/bin/${CROSS}-gcc \
    CPP=/usr/bin/${CROSS}-cpp \
    CXX=/usr/bin/${CROSS}-g++ \
    LD=/usr/bin/${CROSS}-ld
# All packages installed after this we want arm architecture also.
RUN echo "\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty main restricted\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty-updates main restricted\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty universe\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty-updates universe\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty-security main restricted\n\
deb [arch=i386,amd64] http://archive.ubuntu.com/ubuntu/ trusty-security universe\n\
deb [arch=armhf] http://ports.ubuntu.com/ubuntu-ports/ trusty main restricted universe multiverse\n\
deb [arch=armhf] http://ports.ubuntu.com/ubuntu-ports/ trusty-updates main restricted universe multiverse\n\
deb [arch=armhf] http://ports.ubuntu.com/ubuntu-ports/ trusty-security main restricted universe multiverse\n\
deb [arch=armhf] http://ports.ubuntu.com/ubuntu-ports/ trusty-backports main restricted universe multiverse\n\
deb [arch=armhf] http://ports.ubuntu.com/ubuntu-ports/ trusty-proposed main restricted universe multiverse\n"\
> /etc/apt/sources.list && \
  dpkg --add-architecture armhf && \
  apt-get update && apt-get clean && apt-get autoremove --purge
############################
# Golang
############################
ENV GOVERSION go1.8
# Install Golang amd64 
RUN wget https://storage.googleapis.com/golang/${GOVERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf ${GOVERSION}.linux-amd64.tar.gz && \
    rm ${GOVERSION}.linux-amd64.tar.gz 
# Install Golang armv6l 
RUN wget https://storage.googleapis.com/golang/${GOVERSION}.linux-armv6l.tar.gz && \
    tar -xzf ${GOVERSION}.linux-armv6l.tar.gz && \
    cp -R go/pkg/linux_arm /usr/local/go/pkg/ && \
    rm -fr go && rm -frv ${GOVERSION}.linux-armv6l.tar.gz
# Configure Golang
ENV GOPATH=/go \
    GOOS=linux \
    GOARCH=arm \
    GOARM=7 \
    CGO_ENABLED=1
ENV PATH=${PATH}:${GOPATH}/bin:/usr/local/go/bin \
