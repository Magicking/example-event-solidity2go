FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/example-event-solidity2go

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD examples /go/src/${PROJECT_DIR}/examples

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/pongWatcher /go/src/${PROJECT_DIR}/cmd/pongWatcher/main.go

ENTRYPOINT /go/bin/pongWatcher
