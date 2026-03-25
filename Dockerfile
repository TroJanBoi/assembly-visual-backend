FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 9090

CMD ["./main"]
