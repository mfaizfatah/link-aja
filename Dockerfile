FROM golang:1.14.0-alpine3.11 AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/link-aja/

COPY . .

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/link-aja

FROM alpine:3.11

RUN apk add --no-cache tzdata

COPY --from=builder /go/bin/link-aja /go/bin/link-aja

ENTRYPOINT ["/go/bin/link-aja"]