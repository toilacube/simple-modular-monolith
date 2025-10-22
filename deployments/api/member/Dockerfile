# Build stage
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY migrations/ ./migrations/
COPY config.*.yaml ./
COPY docs/ ./docs/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o member_server ./cmd/member_server/main.go

# --- Final stage ---
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/member_server .
COPY --from=builder /app/config.*.yaml ./
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/docs/member ./docs

EXPOSE 6336
CMD ["./member_server"]