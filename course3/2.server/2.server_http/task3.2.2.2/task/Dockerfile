FROM golang:1.19-alpine3.18 AS builder
RUN mkdir ./app
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main

FROM alpine:3.18
WORKDIR /app
COPY --from=builder ./app/.env .
COPY --from=builder ./app/main .
EXPOSE 8000
EXPOSE 8080
CMD ["./main"]
