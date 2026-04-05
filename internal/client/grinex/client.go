package grinex

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
)

// Client loads and calculates spot values from Grinex depth API.
type Client struct {
	httpClient *resty.Client
	baseURL    string
	symbol     string
}

type spotDepthResponse struct {
	Asks json.RawMessage `json:"asks"`
	Bids json.RawMessage `json:"bids"`
}

type depthLevel struct {
	Price string `json:"price"`
}

// New creates a new Grinex API client.
func New(baseURL, symbol string) *Client {
	hc := resty.New()
	hc.SetBaseURL(baseURL)
	return &Client{
		httpClient: hc,
		baseURL:    baseURL,
		symbol:     symbol,
	}
}

// FetchDepth loads depth data from Grinex.
func (c *Client) FetchDepth(ctx context.Context) ([]float64, []float64, error) {
	ctx, span := otel.Tracer("grinex-client").Start(ctx, "FetchDepth")
	defer span.End()

	resp, err := c.httpClient.R().
		SetContext(ctx).
		SetQueryParam("symbol", c.symbol).
		Get("")
	if err != nil {
		return nil, nil, fmt.Errorf("request grinex depth: %w", err)
	}
	if resp.IsError() {
		return nil, nil, fmt.Errorf("grinex returned non-200 status: %s", resp.Status())
	}

	var payload spotDepthResponse
	if err := json.Unmarshal(resp.Body(), &payload); err != nil {
		return nil, nil, fmt.Errorf("decode grinex response: %w", err)
	}

	asks, err := extractPrices(payload.Asks)
	if err != nil {
		return nil, nil, fmt.Errorf("extract asks: %w", err)
	}
	bids, err := extractPrices(payload.Bids)
	if err != nil {
		return nil, nil, fmt.Errorf("extract bids: %w", err)
	}

	return asks, bids, nil
}

func extractPrices(levelsRaw json.RawMessage) ([]float64, error) {
	var levels []json.RawMessage
	if err := json.Unmarshal(levelsRaw, &levels); err != nil {
		return nil, fmt.Errorf("decode levels array: %w", err)
	}

	prices := make([]float64, 0, len(levels))
	for i, raw := range levels {
		price, err := parsePrice(raw)
		if err != nil {
			return nil, fmt.Errorf("parse price at level %d: %w", i, err)
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func parsePrice(raw json.RawMessage) (float64, error) {
	var obj depthLevel
	if err := json.Unmarshal(raw, &obj); err == nil && strings.TrimSpace(obj.Price) != "" {
		price, err := strconv.ParseFloat(obj.Price, 64)
		if err != nil {
			return 0, fmt.Errorf("parse object price: %w", err)
		}
		return price, nil
	}

	var tuple []string
	if err := json.Unmarshal(raw, &tuple); err == nil && len(tuple) > 0 {
		price, err := strconv.ParseFloat(tuple[0], 64)
		if err != nil {
			return 0, fmt.Errorf("parse tuple price: %w", err)
		}
		return price, nil
	}

	return 0, fmt.Errorf("unsupported level format")
}
