FROM vincross/xcompile
# Tools needed by mindcli
RUN apt-get update && \
    apt-get install -y nmap zip curl bc && \
    apt-get clean && apt-get autoremove --purge
# Get wsta for websocket communication with robot.
RUN wget https://github.com/esphen/wsta/releases/download/0.5.0/wsta-0.5.0-x86_64-unknown-linux-gnu.tar.gz && \
    tar -xzvf wsta-0.5.0-x86_64-unknown-linux-gnu.tar.gz -C /usr/local/bin && \
    rm -fr wsta-0.5.0-x86_64-unknown-linux-gnu.tar.gz
RUN wget https://github.com/mholt/caddy/releases/download/v0.9.5/caddy_linux_amd64.tar.gz && \
    tar -xzvf caddy_linux_amd64.tar.gz -C /usr/local/bin && mv /usr/local/bin/caddy_linux_amd64 /usr/local/bin/caddy && \
    rm -fr caddy_linux_amd64.tar.gz
# Add MIND SDK
RUN mkdir -p /.go
ENV MIND_VERSION 0.6.1
RUN wget https://cdn-static.vincross.com/downloads/mind/${MIND_VERSION}/mind.tar.gz
RUN tar -xzf mind.tar.gz -C /.go && rm mind.tar.gz
# Add boilerplate code for `mind init`
ADD boilerplate /.go/src/boilerplate
# Create folder for main entrypoint
RUN mkdir -p /.go/src/skillexec
# Make sure both SDK and Skill is in PATH
ENV GOPATH=/.go:/go
# Add mindcli scripts to PATH
ADD scripts/* /usr/local/mindcli/bin/
ENV PATH=${PATH}:/usr/local/mindcli/bin
# Go to default mount point of Skill
WORKDIR /go/src/skill
