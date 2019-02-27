FROM ubuntu:16.04

# Update packages
RUN apt-get -y update

# Installing useful packages
RUN apt-get install -y wget curl git python tree build-essential apt-utils

# Installing Golang
RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz && mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

USER root

EXPOSE 5000

WORKDIR $GOPATH/src/github.com/FollowInc/Follow_API
ADD . $GOPATH/src/github.com/FollowInc/Follow_API

RUN tree -L 3 ./

RUN go get github.com/gorilla/mux

RUN go build -o api ./src/main.go

# Main command
ENTRYPOINT ["./api"]
CMD [""]