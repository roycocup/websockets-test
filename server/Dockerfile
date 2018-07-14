FROM golang:1.11beta1-alpine3.8

RUN apk update && \
    apk add vim \
    bash \
    git \
    bzr

ENV GOBIN=$GOPATH/bin

COPY src src 

RUN cd src && \ 
    go get ./...

RUN go build -o server src/main.go

EXPOSE 3000

CMD ./server
