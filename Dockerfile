FROM golang:tip-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/api

EXPOSE 8080
CMD ["/server"]
