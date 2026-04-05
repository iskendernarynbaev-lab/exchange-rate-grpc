package config

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "configs/config.yaml"

var envPattern = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)(?::([^}]*))?\}`)

type Config struct {
	GRPCAddr        string
	MetricsAddr     string
	ShutdownTimeout time.Duration

	DatabaseURL string

	GrinexURL string
	Symbol    string

	CalcMethod string
	CalcN      int
	CalcM      int

	OTelServiceName string
	LogLevel        string
}

type fileConfig struct {
	Server struct {
		GRPCAddr        string        `yaml:"grpc_addr"`
		MetricsAddr     string        `yaml:"metrics_addr"`
		ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	} `yaml:"server"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	Grinex struct {
		URL    string `yaml:"url"`
		Symbol string `yaml:"symbol"`
	} `yaml:"grinex"`
	Calculation struct {
		Method string `yaml:"method"`
		N      int    `yaml:"n"`
		M      int    `yaml:"m"`
	} `yaml:"calculation"`
	Observability struct {
		OTelServiceName string `yaml:"otel_service_name"`
		LogLevel        string `yaml:"log_level"`
	} `yaml:"observability"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfgPath := resolveConfigPath()
	cfg, err := loadFromYAML(cfgPath)
	if err != nil {
		return nil, err
	}

	flag.StringVar(&cfg.GRPCAddr, "grpc-addr", cfg.GRPCAddr, "gRPC listen address")
	flag.StringVar(&cfg.MetricsAddr, "metrics-addr", cfg.MetricsAddr, "metrics HTTP listen address")
	flag.DurationVar(&cfg.ShutdownTimeout, "shutdown-timeout", cfg.ShutdownTimeout, "graceful shutdown timeout")
	flag.StringVar(&cfg.DatabaseURL, "database-url", cfg.DatabaseURL, "postgres connection URL")
	flag.StringVar(&cfg.GrinexURL, "grinex-url", cfg.GrinexURL, "grinex depth endpoint")
	flag.StringVar(&cfg.Symbol, "symbol", cfg.Symbol, "grinex symbol")
	flag.StringVar(&cfg.CalcMethod, "calc-method", cfg.CalcMethod, "calculation method: topn or avgnm")
	flag.IntVar(&cfg.CalcN, "calc-n", cfg.CalcN, "N index for topN or lower range for avgNM")
	flag.IntVar(&cfg.CalcM, "calc-m", cfg.CalcM, "M upper range for avgNM")
	flag.StringVar(&cfg.OTelServiceName, "otel-service-name", cfg.OTelServiceName, "OpenTelemetry service name")
	flag.StringVar(&cfg.LogLevel, "log-level", cfg.LogLevel, "log level: debug, info, warn, error")
	flag.StringVar(&cfgPath, "config", cfgPath, "path to YAML config file")
	flag.Parse()

	cfg.CalcMethod = strings.ToLower(cfg.CalcMethod)
	cfg.LogLevel = strings.ToLower(cfg.LogLevel)

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadFromYAML(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file %q: %w", path, err)
	}

	expanded := expandEnv(string(content))

	var raw fileConfig
	if err := yaml.Unmarshal([]byte(expanded), &raw); err != nil {
		return nil, fmt.Errorf("parse config file %q: %w", path, err)
	}

	cfg := &Config{
		GRPCAddr:        raw.Server.GRPCAddr,
		MetricsAddr:     raw.Server.MetricsAddr,
		ShutdownTimeout: raw.Server.ShutdownTimeout,
		DatabaseURL:     raw.Database.URL,
		GrinexURL:       raw.Grinex.URL,
		Symbol:          raw.Grinex.Symbol,
		CalcMethod:      raw.Calculation.Method,
		CalcN:           raw.Calculation.N,
		CalcM:           raw.Calculation.M,
		OTelServiceName: raw.Observability.OTelServiceName,
		LogLevel:        raw.Observability.LogLevel,
	}

	return cfg, nil
}

func resolveConfigPath() string {
	path := strings.TrimSpace(os.Getenv("CONFIG_PATH"))
	if path == "" {
		path = defaultConfigPath
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "-config" && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
		if strings.HasPrefix(arg, "-config=") {
			return strings.TrimPrefix(arg, "-config=")
		}
	}

	return path
}

func expandEnv(in string) string {
	return envPattern.ReplaceAllStringFunc(in, func(match string) string {
		parts := envPattern.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}

		name := parts[1]
		def := ""
		if len(parts) > 2 {
			def = parts[2]
		}

		value, ok := os.LookupEnv(name)
		if ok && strings.TrimSpace(value) != "" {
			return value
		}
		return def
	})
}

func (c *Config) Validate() error {
	if c.GRPCAddr == "" {
		return fmt.Errorf("grpc addr must not be empty")
	}
	if c.DatabaseURL == "" {
		return fmt.Errorf("database url must not be empty")
	}
	if c.GrinexURL == "" {
		return fmt.Errorf("grinex url must not be empty")
	}
	if c.Symbol == "" {
		return fmt.Errorf("symbol must not be empty")
	}
	if c.CalcN <= 0 {
		return fmt.Errorf("calc N must be greater than 0")
	}
	if c.CalcMethod != "topn" && c.CalcMethod != "avgnm" {
		return fmt.Errorf("calc method must be topn or avgnm")
	}
	if c.CalcMethod == "avgnm" && c.CalcM < c.CalcN {
		return fmt.Errorf("calc M must be greater than or equal to calc N for avgnm")
	}
	if c.ShutdownTimeout <= 0 {
		return fmt.Errorf("shutdown timeout must be greater than 0")
	}
	return nil
}
