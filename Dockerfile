FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o static ./cmd/static/static.go

EXPOSE 8081

CMD ["./static"]
