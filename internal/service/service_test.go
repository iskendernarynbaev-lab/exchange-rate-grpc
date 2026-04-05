package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/metrics"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/model"
)

type mockClient struct {
	asks []float64
	bids []float64
	err  error
}

func (m mockClient) FetchDepth(context.Context) ([]float64, []float64, error) {
	if m.err != nil {
		return nil, nil, m.err
	}
	return m.asks, m.bids, nil
}

type mockRepo struct {
	stored []model.Rate
	err    error
}

func (m *mockRepo) StoreRate(_ context.Context, rate model.Rate) error {
	if m.err != nil {
		return m.err
	}
	m.stored = append(m.stored, rate)
	return nil
}

func TestServiceGetRatesSuccess(t *testing.T) {
	reg := prometheus.NewRegistry()
	met := metrics.New(reg)
	repo := &mockRepo{}
	svc := New(
		mockClient{asks: []float64{100, 101}, bids: []float64{99, 98}},
		repo,
		met,
		"topn",
		1,
		2,
	)

	rate, err := svc.GetRates(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rate.Ask != 100 {
		t.Fatalf("unexpected ask: got %v", rate.Ask)
	}
	if rate.Bid != 99 {
		t.Fatalf("unexpected bid: got %v", rate.Bid)
	}
	if time.Since(rate.CreatedAt) > 2*time.Second {
		t.Fatalf("unexpected timestamp: %v", rate.CreatedAt)
	}
	if len(repo.stored) != 1 {
		t.Fatalf("rate not stored")
	}
}

func TestServiceGetRatesFetchError(t *testing.T) {
	reg := prometheus.NewRegistry()
	met := metrics.New(reg)
	svc := New(mockClient{err: errors.New("boom")}, &mockRepo{}, met, "topn", 1, 1)

	_, err := svc.GetRates(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestServiceGetRatesStoreError(t *testing.T) {
	reg := prometheus.NewRegistry()
	met := metrics.New(reg)
	svc := New(
		mockClient{asks: []float64{100}, bids: []float64{99}},
		&mockRepo{err: errors.New("db error")},
		met,
		"topn",
		1,
		1,
	)

	_, err := svc.GetRates(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
}
