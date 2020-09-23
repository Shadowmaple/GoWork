package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Error("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", "localhost:88"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
