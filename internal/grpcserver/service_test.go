package grpcserver

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/metrics"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/model"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/service"
)

type grpcMockClient struct{}

func (grpcMockClient) FetchDepth(context.Context) ([]float64, []float64, error) {
	return []float64{100.5}, []float64{99.5}, nil
}

type grpcMockRepo struct{}

func (grpcMockRepo) StoreRate(context.Context, model.Rate) error { return nil }

func TestRatesServerGetRates(t *testing.T) {
	reg := prometheus.NewRegistry()
	met := metrics.New(reg)
	svc := service.New(grpcMockClient{}, grpcMockRepo{}, met, "topn", 1, 1)
	srv := NewRatesServer(svc, zap.NewNop())

	resp, err := srv.GetRates(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.GetAsk() == 0 {
		t.Fatal("ask field missing")
	}
	if resp.GetBid() == 0 {
		t.Fatal("bid field missing")
	}
	if resp.GetTimestamp() == "" {
		t.Fatal("timestamp field missing")
	}
}
