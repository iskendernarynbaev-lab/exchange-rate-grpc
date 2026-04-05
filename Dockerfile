FROM golang:1.25.7 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/exchange-rate-grpc ./cmd/exchange-rate-grpc

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/exchange-rate-grpc /app/exchange-rate-grpc
COPY --from=builder /src/configs /app/configs
EXPOSE 9090 2112
ENTRYPOINT ["/app/exchange-rate-grpc"]
