FROM golang:1.21.6-alpine as builder

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go build -o /bin/web ./cmd/web/main.go

FROM alpine:latest as production

COPY --from=builder /bin/web /bin/web

RUN chmod +x /bin/web
