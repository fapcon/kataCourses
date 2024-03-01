FROM golang:1.19
RUN mkdir ./app
COPY .. ./app
WORKDIR ./app
RUN go build -o main
EXPOSE 8080
CMD ["./main"]
