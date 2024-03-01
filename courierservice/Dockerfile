FROM golang:1.19.1-alpine3.15 AS builder
COPY . /go/src/courierservice
WORKDIR /go/src/courierservice
# Create slimest possible image
RUN go build -ldflags="-w -s" -o /go/bin/geogrpc /go/src/courierservice/cmd/api

FROM alpine:3.15
# Copy binary from builder
COPY --from=builder /go/bin/server /go/bin/server
COPY ./public /app/public
COPY ./.env /app/.env

WORKDIR /app
# Set entrypoint
ENTRYPOINT ["/go/bin/server"]