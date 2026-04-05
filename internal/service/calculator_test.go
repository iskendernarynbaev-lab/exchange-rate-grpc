package service

import "testing"

func TestCalculateTopN(t *testing.T) {
	got, err := Calculate("topn", 2, 0, []float64{10, 20, 30})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 20 {
		t.Fatalf("unexpected value: got %v want 20", got)
	}
}

func TestCalculateAvgNM(t *testing.T) {
	got, err := Calculate("avgnm", 2, 4, []float64{10, 20, 30, 40, 50})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 30 {
		t.Fatalf("unexpected value: got %v want 30", got)
	}
}

func TestCalculateUnknownMethod(t *testing.T) {
	_, err := Calculate("invalid", 1, 1, []float64{10})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCalculateTopNRangeError(t *testing.T) {
	_, err := Calculate("topn", 3, 0, []float64{10, 20})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCalculateAvgNMRangeError(t *testing.T) {
	_, err := Calculate("avgnm", 3, 2, []float64{10, 20, 30})
	if err == nil {
		t.Fatal("expected error")
	}
}
