FROM golang:1.7.3-wheezy

ADD . /go/src/github.com/transitorykris/hypnos
RUN cd /go/src/github.com/transitorykris/hypnos && go get && go build -o /hypnos
