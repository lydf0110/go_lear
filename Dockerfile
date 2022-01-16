FROM  golang:latest  as builder
MAINTAINER  jackie
WORKDIR  /go/src
COPY  .  .
RUN  go  build  -o  go_lear

FROM  alpine:latest
WORKDIR  /root
COPY  --from=builder /go/src/go_lear  .
EXPOSE 8080
CMD  [“./go_lear”]
