FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/andrewbreksa/httpforwarder/

COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /httpforwarder

FROM scratch
COPY --from=builder /httpforwarder /httpforwarder
ENTRYPOINT ["/httpforwarder"]