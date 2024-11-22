package prettylogger

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"sync"
	"task01/pkg/prettylogger/handlers"
)

func NewColoredHandler(w io.Writer, opts *slog.HandlerOptions) *handlers.TextHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	b := &bytes.Buffer{}
	return &handlers.TextHandler{
		H: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: handlers.SupressDefaults(opts.ReplaceAttr),
		}),
		B:    b,
		M:    &sync.Mutex{},
		Opts: opts,
		W:    w,
	}
}

func NewJsonHandler(w io.Writer, opts *slog.HandlerOptions) *handlers.JsonHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	b := &bytes.Buffer{}
	return &handlers.JsonHandler{
		H: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: opts.ReplaceAttr,
		}),
		B:    b,
		M:    &sync.Mutex{},
		Opts: opts,
		W:    w,
	}
}

func NewDiscardHandler() *handlers.DiscardHandler {
	return &handlers.DiscardHandler{}
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func New(debugLevel string) *slog.Logger {
	logger := slog.New(NewColoredHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}))
	switch debugLevel {
	case "dev":
		{
			logger = slog.New(NewJsonHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}))
		}
	case "local":
		{
			logger = slog.New(NewColoredHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}))
		}
	case "prod":
		{
			logger = slog.New(NewJsonHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelInfo,
				ReplaceAttr: nil,
			}))
		}

	}
	return logger
}
