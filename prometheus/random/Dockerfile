FROM golang:1.9.3 AS builder
WORKDIR /go/src/github.com/prometheus
RUN go get -d github.com/prometheus/client_golang/examples/random
WORKDIR /go/src/github.com/prometheus/client_golang/examples/random
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

# Final image.
FROM scratch
COPY --from=builder /go/src/github.com/prometheus/client_golang/examples/random .
EXPOSE 8080
ENTRYPOINT ["/random"]
