FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o logger-service ./cmd/api

FROM scratch
COPY --from=builder /app .
CMD [ "./logger-service" ]