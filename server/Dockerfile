ARG GO_VERSION=1.16

FROM golang:${GO_VERSION} as build
WORKDIR /app
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server .

FROM alpine:latest
LABEL org.opencontainers.image.source = "https://github.com/reliablyhq/walkthrough-slo-http-server"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/bin/server .
CMD ["./server"]