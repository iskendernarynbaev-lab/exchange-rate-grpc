package health

import (
	"google.golang.org/grpc/health"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

func New() *health.Server {
	h := health.NewServer()
	h.SetServingStatus("", grpcHealthV1.HealthCheckResponse_SERVING)
	return h
}
