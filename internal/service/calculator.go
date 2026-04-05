package service

import "fmt"

func Calculate(method string, n, m int, values []float64) (float64, error) {
	switch method {
	case "topn":
		return topN(values, n)
	case "avgnm":
		return avgNM(values, n, m)
	default:
		return 0, fmt.Errorf("unknown method %q", method)
	}
}

func topN(values []float64, n int) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("N must be greater than 0")
	}
	if len(values) < n {
		return 0, fmt.Errorf("insufficient values: want index %d, got %d", n, len(values))
	}
	return values[n-1], nil
}

func avgNM(values []float64, n, m int) (float64, error) {
	if n <= 0 || m <= 0 {
		return 0, fmt.Errorf("N and M must be greater than 0")
	}
	if m < n {
		return 0, fmt.Errorf("M must be greater than or equal to N")
	}
	if len(values) < m {
		return 0, fmt.Errorf("insufficient values: want M=%d, got %d", m, len(values))
	}

	var sum float64
	for i := n - 1; i <= m-1; i++ {
		sum += values[i]
	}
	count := float64(m - n + 1)
	return sum / count, nil
}
