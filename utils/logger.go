package utils

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func SetLogger() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	Logger = slog.New(
		slog.NewTextHandler(os.Stdout, opts),
	)
}
