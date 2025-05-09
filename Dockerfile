FROM golang:1.23-alpine
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/chat-app github.com/AliMumtazDev/Go_Chat_App
RUN chmod +x /app/chat-app
EXPOSE 8002
CMD ["/app/chat-app"]