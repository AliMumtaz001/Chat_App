FROM golang:1.16
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o websocket-server ./main.go
EXPOSE 8004
CMD ["./websocket-server"]