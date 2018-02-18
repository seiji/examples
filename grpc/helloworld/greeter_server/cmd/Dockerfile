FROM golang:1.9.3 AS builder
RUN go get -d github.com/seiji/examples/grpc/helloworld/greeter_server
WORKDIR /go/src/github.com/seiji/examples/grpc/helloworld/greeter_server
RUN go get -u github.com/golang/dep/cmd/dep
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'
RUN ls -la

FROM scratch
COPY --from=builder /go/src/github.com/seiji/examples/grpc/helloworld/greeter_server/greeter_server .
EXPOSE 50051

ENTRYPOINT ["./greeter_server"]
