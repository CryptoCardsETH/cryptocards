FROM php:7.1-fpm

RUN apt-get update
RUN apt-get install -y git libgflags-dev build-essential automake autoconf libtool shtool curl unzip
WORKDIR /tmp
RUN git clone https://github.com/grpc/grpc

WORKDIR /tmp/grpc
RUN git submodule update --init

RUN make grpc_php_plugin

WORKDIR /tmp

RUN curl -L -o protoc_binaries.zip "https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip"
RUN unzip protoc_binaries.zip -d "proto_tools"



# ADD docker-run.sh /docker-run.sh

ENTRYPOINT ["/bin/bash","/tools/docker_run.sh"] 
CMD ["run"]

