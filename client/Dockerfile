# Build stage
FROM golang:1.11-alpine as builder

RUN apk update

WORKDIR /go/src/github.com/deciphernow/fail/client
ADD . /go/src/github.com/deciphernow/fail/client

# Compile necessary binaries for final image
RUN go build

# Run stage
FROM alpine:3.7

# Necessary for cloudwatch put calls to AWS
RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

# Copy over gm-proxy binary
COPY --from=builder /go/src/github.com/deciphernow/fail/client/client /app/

EXPOSE 8080

CMD ./client
