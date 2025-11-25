FROM golang:latest

WORKDIR /usr/local/app
COPY . .
EXPOSE 8080

RUN go build -o compute-max-server

CMD ["./compute-max-server"]
