package grpcserver

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/service"
	ratesv1 "github.com/iskendernarynbaev-lab/exchange-rate-grpc/pkg/api/rates/v1"
)

type RatesServer struct {
	svc    *service.Service
	logger *zap.Logger
	ratesv1.UnimplementedRatesServiceServer
}

func NewRatesServer(svc *service.Service, logger *zap.Logger) *RatesServer {
	return &RatesServer{svc: svc, logger: logger}
}

func (s *RatesServer) GetRates(ctx context.Context, _ *emptypb.Empty) (*ratesv1.GetRatesResponse, error) {
	start := time.Now()

	rate, err := s.svc.GetRates(ctx)
	if err != nil {
		s.logger.Warn("GetRates failed",
			zap.Error(err),
			zap.Duration("duration", time.Since(start)),
		)
		return nil, status.Errorf(codes.Internal, "get rates: %v", err)
	}

	s.logger.Info("GetRates completed",
		zap.Float64("ask", rate.Ask),
		zap.Float64("bid", rate.Bid),
		zap.String("timestamp", rate.CreatedAt.Format(time.RFC3339Nano)),
		zap.Duration("duration", time.Since(start)),
	)

	return &ratesv1.GetRatesResponse{
		Ask:        rate.Ask,
		Bid:        rate.Bid,
		Timestamp:  rate.CreatedAt.Format("2006-01-02T15:04:05.999999999Z07:00"),
		UnixMillis: rate.CreatedAt.UnixMilli(),
	}, nil
}
