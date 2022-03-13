FROM  golang:latest  as builder
MAINTAINER  jackie
WORKDIR  /go/src
COPY  .  .
RUN  go  build  -o  go_lear
EXPOSE 8080
CMD  ["./go/src/go_lear"]
