FROM golang:1.10.0 AS builder
RUN go get -d github.com/seiji/examples/grpc/helloworld/greeter_server
WORKDIR /go/src/github.com/seiji/examples/grpc/helloworld/greeter_server
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' cmd/*.go
RUN ls -la

FROM scratch
COPY --from=builder /go/src/github.com/seiji/examples/grpc/helloworld/greeter_server/main .
EXPOSE 50051

ENTRYPOINT ["./main"]
