# Stage 1: Building the code
FROM golang:1.18-buster as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tinder ./cmd/person

# Stage 2: Setup runtime container
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tinder .
EXPOSE 8080
CMD ["./tinder"]
