FROM golang:wheezy

WORKDIR /go
ADD . /go

RUN echo $GOPATH

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN cd src/GoRpc && dep ensure -v -vendor-only

ENTRYPOINT ["/bin/bash","/go/docker_run.sh"] 
CMD ["run"]
