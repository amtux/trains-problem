FROM golang:1.11-stretch

WORKDIR /go/src/trains-problem/

COPY . .

RUN make
