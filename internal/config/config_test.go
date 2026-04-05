package config

import (
	"os"
	"testing"
)

func TestExpandEnvWithDefault(t *testing.T) {
	t.Setenv("TEST_KEY", "")
	got := expandEnv("value=${TEST_KEY:default}")
	if got != "value=default" {
		t.Fatalf("unexpected value: %s", got)
	}
}

func TestExpandEnvWithActualValue(t *testing.T) {
	t.Setenv("TEST_KEY", "real")
	got := expandEnv("value=${TEST_KEY:default}")
	if got != "value=real" {
		t.Fatalf("unexpected value: %s", got)
	}
}

func TestLoadFromYAML(t *testing.T) {
	t.Setenv("GRPC_ADDR", "")
	t.Setenv("METRICS_ADDR", "")
	t.Setenv("SHUTDOWN_TIMEOUT", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("GRINEX_URL", "")
	t.Setenv("GRINEX_SYMBOL", "")
	t.Setenv("CALC_METHOD", "")
	t.Setenv("CALC_N", "")
	t.Setenv("CALC_M", "")
	t.Setenv("OTEL_SERVICE_NAME", "")
	t.Setenv("LOG_LEVEL", "")

	content := []byte(`
server:
  grpc_addr: ${GRPC_ADDR::9000}
  metrics_addr: ${METRICS_ADDR::9200}
  shutdown_timeout: ${SHUTDOWN_TIMEOUT:8s}
database:
  url: ${DATABASE_URL:postgres://x}
grinex:
  url: ${GRINEX_URL:https://example.com}
  symbol: ${GRINEX_SYMBOL:test}
calculation:
  method: ${CALC_METHOD:topn}
  n: ${CALC_N:2}
  m: ${CALC_M:4}
observability:
  otel_service_name: ${OTEL_SERVICE_NAME:test-service}
  log_level: ${LOG_LEVEL:debug}
`)

	tmp, err := os.CreateTemp(t.TempDir(), "cfg-*.yaml")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	defer tmp.Close()

	if _, err := tmp.Write(content); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	cfg, err := loadFromYAML(tmp.Name())
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.GRPCAddr != ":9000" {
		t.Fatalf("unexpected grpc addr: %s", cfg.GRPCAddr)
	}
	if cfg.CalcN != 2 || cfg.CalcM != 4 {
		t.Fatalf("unexpected calc range: %d-%d", cfg.CalcN, cfg.CalcM)
	}
}
