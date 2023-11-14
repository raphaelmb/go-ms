FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o broker-service ./cmd/api

FROM scratch
COPY --from=builder /app .
CMD [ "./broker-service" ]