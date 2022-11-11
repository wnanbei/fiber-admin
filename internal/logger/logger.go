package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

func init() {
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.DebugLevel,
		ReplaceAttr: nil,
	}

	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stderr)))
}
