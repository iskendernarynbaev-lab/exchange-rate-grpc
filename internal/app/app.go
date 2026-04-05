package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/client/grinex"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/config"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/grpcserver"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/health"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/logger"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/metrics"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/repository/postgres"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/service"
	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/telemetry"
	ratesv1 "github.com/iskendernarynbaev-lab/exchange-rate-grpc/pkg/api/rates/v1"
)

func Run() error {
	startedAt := time.Now()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	log, err := logger.New(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}
	defer func() { _ = log.Sync() }()

	log.Info("application starting",
		zap.String("grpc_addr", cfg.GRPCAddr),
		zap.String("metrics_addr", cfg.MetricsAddr),
		zap.String("symbol", cfg.Symbol),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	shutdownTracer, err := telemetry.Init(ctx, cfg.OTelServiceName)
	if err != nil {
		return fmt.Errorf("init telemetry: %w", err)
	}
	log.Info("telemetry initialized", zap.String("service", cfg.OTelServiceName))

	repo, err := postgres.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("init repository: %w", err)
	}
	log.Info("database connection established")
	defer func() {
		if err := repo.Close(); err != nil {
			log.Error("repository close failed", zap.Error(err))
		}
	}()

	registry := prometheus.NewRegistry()
	m := metrics.New(registry)

	marketClient := grinex.New(cfg.GrinexURL, cfg.Symbol)
	svc := service.New(marketClient, repo, m)
	ratesServer := grpcserver.NewRatesServer(svc, log)
	healthServer := health.New()
	log.Info("service dependencies initialized")

	grpcSrv := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	ratesv1.RegisterRatesServiceServer(grpcSrv, ratesServer)
	grpcHealthV1.RegisterHealthServer(grpcSrv, healthServer)
	reflection.Register(grpcSrv)

	grpcLn, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return fmt.Errorf("listen gRPC: %w", err)
	}

	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	metricsMux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	metricsSrv := &http.Server{
		Addr:              cfg.MetricsAddr,
		Handler:           otelhttp.NewHandler(metricsMux, "metrics-http"),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 2)
	go func() {
		log.Info("gRPC server started", zap.String("addr", cfg.GRPCAddr))
		if err := grpcSrv.Serve(grpcLn); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			errCh <- fmt.Errorf("serve gRPC: %w", err)
		}
	}()
	go func() {
		log.Info("metrics server started", zap.String("addr", cfg.MetricsAddr))
		if err := metricsSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("serve metrics: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		log.Info("shutdown signal received")
	case err := <-errCh:
		cancel()
		log.Error("server failure", zap.Error(err))
	}

	log.Info("shutdown started", zap.Duration("timeout", cfg.ShutdownTimeout))

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer shutdownCancel()

	gracefulDone := make(chan struct{})
	go func() {
		grpcSrv.GracefulStop()
		close(gracefulDone)
	}()

	select {
	case <-gracefulDone:
		log.Info("gRPC server stopped gracefully")
	case <-shutdownCtx.Done():
		grpcSrv.Stop()
		log.Warn("gRPC force stop applied")
	}

	if err := metricsSrv.Shutdown(shutdownCtx); err != nil {
		log.Error("metrics shutdown failed", zap.Error(err))
	}

	if err := shutdownTracer(shutdownCtx); err != nil {
		log.Error("telemetry shutdown failed", zap.Error(err))
	}

	log.Info("application stopped", zap.Duration("uptime", time.Since(startedAt)))

	return nil
}
