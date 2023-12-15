# Builder stage
FROM golang:latest as builder
WORKDIR /app

# Kopyala ve bağımlılıkları çöz
COPY go.mod go.sum ./
RUN go mod download

# Uygulama kodunu kopyala
COPY . .

# Uygulamayı derle
RUN go build -o my-app .

# Runtime stage
FROM golang:latest

WORKDIR /app

# Builder aşamasında derlenen uygulamayı kopyala
COPY --from=builder /app/my-app .
COPY .env .
# Çalıştır
CMD ["./my-app"]
