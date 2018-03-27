FROM golang:wheezy

WORKDIR /go
ADD . /go

# install dep, install the deps as per Gopkg.lock
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd src/GoRpc && dep ensure -v -vendor-only

# run the rpc server
ENTRYPOINT ["/bin/bash","/go/docker_run.sh"] 
CMD ["run"]
