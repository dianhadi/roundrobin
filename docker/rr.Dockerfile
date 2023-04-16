FROM golang:alpine

WORKDIR /dir

COPY . .

RUN go build -o rr ./cmd/rr

EXPOSE 8080

CMD ["./rr"]
