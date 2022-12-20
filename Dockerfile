FROM golang:1.19.4 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
## Build our application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./...

## Use the lightweight scratch image to
## run our application within
FROM alpine:latest AS production
## Add debuging tools
RUN apk add postgresql14-client curl
WORKDIR /app
## Copy binary from builder stage
COPY --from=builder /app/lorem-seed .
