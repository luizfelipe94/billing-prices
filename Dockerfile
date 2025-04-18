FROM golang:1.24
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go
EXPOSE 8081
CMD ["./main"]
