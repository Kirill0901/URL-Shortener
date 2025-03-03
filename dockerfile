FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .
COPY . .

RUN go build -o url-shortener /build/cmd/url-shortener/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/url-shortener /build/url-shortener

EXPOSE 12000

ENTRYPOINT ["/build/url-shortener"]
