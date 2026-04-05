package main

import (
	"fmt"
	"os"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fatal: %v\n", err)
		os.Exit(1)
	}
}
