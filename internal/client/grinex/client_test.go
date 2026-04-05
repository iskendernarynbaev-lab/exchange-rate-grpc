package grinex

import (
	"encoding/json"
	"testing"
)

func TestExtractPricesTupleFormat(t *testing.T) {
	raw := json.RawMessage(`[["1.1","100"],["1.2","200"]]`)
	got, err := extractPrices(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0] != 1.1 || got[1] != 1.2 {
		t.Fatalf("unexpected prices: %#v", got)
	}
}

func TestExtractPricesObjectFormat(t *testing.T) {
	raw := json.RawMessage(`[{"price":"80.88","volume":"72421.3081","amount":"5857435.4"},{"price":"80.94","volume":"20000.0","amount":"1618800.0"}]`)
	got, err := extractPrices(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 || got[0] != 80.88 || got[1] != 80.94 {
		t.Fatalf("unexpected prices: %#v", got)
	}
}

func TestExtractPricesInvalid(t *testing.T) {
	raw := json.RawMessage(`[{"price":"bad"}]`)
	_, err := extractPrices(raw)
	if err == nil {
		t.Fatal("expected error")
	}
}
