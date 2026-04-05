package grpcserver

import (
	"context"
	"strings"
	"time"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/service"
	ratesv1 "github.com/iskendernarynbaev-lab/exchange-rate-grpc/pkg/api/rates/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RatesServer struct {
	svc    *service.Service
	logger *zap.Logger
	ratesv1.UnimplementedRatesServiceServer
}

func NewRatesServer(svc *service.Service, logger *zap.Logger) *RatesServer {
	return &RatesServer{svc: svc, logger: logger}
}

func (s *RatesServer) GetRates(ctx context.Context, req *ratesv1.GetRatesRequest) (*ratesv1.GetRatesResponse, error) {
	start := time.Now()
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request must not be nil")
	}
	method := strings.ToLower(strings.TrimSpace(req.GetMethod()))
	n := int(req.GetN())
	m := int(req.GetM())

	if method != "topn" && method != "avgnm" {
		return nil, status.Error(codes.InvalidArgument, "method must be topn or avgnm")
	}
	if n <= 0 {
		return nil, status.Error(codes.InvalidArgument, "n must be greater than 0")
	}
	if method == "avgnm" {
		if m <= 0 {
			return nil, status.Error(codes.InvalidArgument, "m must be greater than 0")
		}
		if m < n {
			return nil, status.Error(codes.InvalidArgument, "m must be greater than or equal to n")
		}
	}

	rate, err := s.svc.GetRates(ctx, method, n, m)
	if err != nil {
		s.logger.Warn("GetRates failed",
			zap.Error(err),
			zap.Duration("duration", time.Since(start)),
			zap.String("method", method),
			zap.Int("n", n),
			zap.Int("m", m),
		)
		return nil, status.Errorf(codes.Internal, "get rates: %v", err)
	}

	s.logger.Info("GetRates completed",
		zap.Float64("ask", rate.Ask),
		zap.Float64("bid", rate.Bid),
		zap.String("timestamp", rate.CreatedAt.Format(time.RFC3339Nano)),
		zap.String("method", method),
		zap.Int("n", n),
		zap.Int("m", m),
		zap.Duration("duration", time.Since(start)),
	)

	return &ratesv1.GetRatesResponse{
		Ask:        rate.Ask,
		Bid:        rate.Bid,
		Timestamp:  rate.CreatedAt.Format("2006-01-02T15:04:05.999999999Z07:00"),
		UnixMillis: rate.CreatedAt.UnixMilli(),
	}, nil
}
