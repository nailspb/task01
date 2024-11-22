package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"task01/pkg/prettylogger/colors"
)

type TextHandler struct {
	H    slog.Handler
	B    *bytes.Buffer
	M    *sync.Mutex
	Opts *slog.HandlerOptions
	W    io.Writer
}

func (h *TextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.Opts.Level.Level()
}

func (h *TextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TextHandler{H: h.H.WithAttrs(attrs), B: h.B, M: h.M, Opts: h.Opts, W: h.W}
}

func (h *TextHandler) WithGroup(name string) slog.Handler {
	return &TextHandler{H: h.H.WithGroup(name), B: h.B, M: h.M, Opts: h.Opts, W: h.W}
}

func (h *TextHandler) computeAttr(ctx context.Context, r slog.Record) (map[string]any, error) {
	h.M.Lock()
	defer func() {
		h.B.Reset()
		h.M.Unlock()
	}()
	if err := h.H.Handle(ctx, r); err != nil {
		return nil, fmt.Errorf("Error when calling inner handler's Handle: %W", err)
	}

	var attrs map[string]any
	err := json.Unmarshal(h.B.Bytes(), &attrs)
	if err != nil {
		return nil, fmt.Errorf("Error when unmarshalling inner handler's Handle attrs: %W", err)
	}
	return attrs, nil
}

func (h *TextHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()

	switch r.Level {
	case slog.LevelDebug:

		level = colors.Cyan(level)
		r.Message = colors.Cyan(r.Message)
	case slog.LevelInfo:
		level = colors.Blue(level)
		r.Message = colors.Blue(r.Message)
	case slog.LevelWarn:
		level = colors.Yellow(level)
		r.Message = colors.Yellow(r.Message)
	case slog.LevelError:
		level = colors.Red(level)
		r.Message = colors.Red(r.Message)
	}

	attrs, err := h.computeAttr(ctx, r)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(attrs, "", " ")
	if err != nil {
		return fmt.Errorf("Error when unmarshaling attrs: %v", err)
	}
	if len(bytes) > 2 {
		fmt.Fprintf(
			h.W,
			"%s %s: %s %s\n",
			colors.White(r.Time.Format("[2006-01-02 15:04:05.000]")),
			level,
			r.Message,
			colors.HiWhite("\n"+string(bytes)),
		)
	} else {
		fmt.Fprintf(
			h.W,
			"%s %s: %s\n",
			colors.White(r.Time.Format("[2006-01-02 15:04:05.000]")),
			level,
			colors.White(r.Message),
		)
	}

	return nil
}
