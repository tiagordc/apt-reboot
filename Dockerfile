# STEP 1 build go binary
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/apt-reboot
COPY src/ src/
COPY go.mod .
RUN go get -d -v ./src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./src

# STEP 2 build a small docker image
FROM scratch
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
