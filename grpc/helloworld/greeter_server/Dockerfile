FROM golang:1.9-alpine

RUN \
  apk --update add --no-cache git &&\
  go get -u github.com/seiji/grpc-examples/helloworld/greeter_server/cmd &&\
  apk del git &&\
  rm -rf /go/src

EXPOSE 50051

CMD ["greeter_server"]
