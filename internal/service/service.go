package service

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/metrics"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/model"
)

type MarketClient interface {
	FetchDepth(ctx context.Context) ([]float64, []float64, error)
}

type RateRepository interface {
	StoreRate(ctx context.Context, rate model.Rate) error
}

type Service struct {
	client  MarketClient
	repo    RateRepository
	metrics *metrics.Metrics
	method  string
	n       int
	m       int
}

func New(client MarketClient, repo RateRepository, metrics *metrics.Metrics, method string, n, m int) *Service {
	return &Service{client: client, repo: repo, metrics: metrics, method: method, n: n, m: m}
}

func (s *Service) GetRates(ctx context.Context) (model.Rate, error) {
	start := time.Now()
	ctx, span := otel.Tracer("rates-service").Start(ctx, "GetRates")
	defer span.End()

	asks, bids, err := s.client.FetchDepth(ctx)
	if err != nil {
		s.observe("error", start)
		return model.Rate{}, fmt.Errorf("fetch market data: %w", err)
	}

	ask, err := Calculate(s.method, s.n, s.m, asks)
	if err != nil {
		s.observe("error", start)
		return model.Rate{}, fmt.Errorf("calculate ask: %w", err)
	}

	bid, err := Calculate(s.method, s.n, s.m, bids)
	if err != nil {
		s.observe("error", start)
		return model.Rate{}, fmt.Errorf("calculate bid: %w", err)
	}

	rate := model.Rate{
		Ask:       ask,
		Bid:       bid,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.StoreRate(ctx, rate); err != nil {
		s.metrics.StoreErrors.Inc()
		s.observe("error", start)
		return model.Rate{}, fmt.Errorf("store rate: %w", err)
	}

	s.observe("ok", start)
	return rate, nil
}

func (s *Service) observe(status string, start time.Time) {
	s.metrics.RequestsTotal.WithLabelValues(status).Inc()
	s.metrics.Duration.WithLabelValues(status).Observe(time.Since(start).Seconds())
}
