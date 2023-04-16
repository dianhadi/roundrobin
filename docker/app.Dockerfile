FROM golang:alpine

WORKDIR /dir

COPY . .

RUN go build -o app ./cmd/app

EXPOSE 8080

CMD ["./app"]
