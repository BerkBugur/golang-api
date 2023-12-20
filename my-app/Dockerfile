# Builder stage
FROM golang:latest as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o my-app .

# Runtime stage
FROM golang:latest

WORKDIR /app

# Copy application compiled in builder stage
COPY --from=builder /app/my-app .
COPY .env .
# RUN
CMD ["./my-app"]
