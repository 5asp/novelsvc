FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.io"

WORKDIR /app 

COPY go.mod .


RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -v -o tonovel-go cmd/novelsvc/main.go

FROM scratch

WORKDIR /app


COPY --from=builder /app/tonovel-go .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


EXPOSE 8081

ENTRYPOINT ["/app/tonovel-go"]
