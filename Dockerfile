ARG GO_VERSION=0

FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o snippetbox ./cmd/web

FROM scratch
LABEL org.opencontainers.image.source="https://github.com/afonsodemori/snippetbox"
LABEL org.opencontainers.image.licenses="MIT"
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/snippetbox .
USER 1000:1000
EXPOSE 4000
ENTRYPOINT ["./snippetbox"]
