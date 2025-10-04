package handlers

import (
	"context"
	"io"
	"log/slog"
	"strings"
	"time"
)

type EmojiHandler struct {
	Out         io.Writer
	LevelEmojis map[slog.Level]string
	TimeFormat  string
	LevelVar    *slog.LevelVar

	attrs  []slog.Attr
	groups []string
}

func (h *EmojiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// Always enabled; user can wrap logger with slog.NewTextHandler with Level
	return true
}

func (h *EmojiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// return a shallow copy with appended attrs
	cp := *h
	cp.attrs = append(cp.attrs[:len(cp.attrs):len(cp.attrs)], attrs...)
	return &cp
}

func (h *EmojiHandler) WithGroup(name string) slog.Handler {
	cp := *h
	cp.groups = append(cp.groups[:len(cp.groups):len(cp.groups)], name)
	return &cp
}

func (h *EmojiHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level < h.LevelVar.Level() {
		return nil
	}

	// Build a list of attributes (attrs from WithAttrs first, then record attrs)
	// We'll inspect these to find "emoji", "log_type", and "no_newline".
	var (
		emoji     string
		logType   string
		noNewline bool
	)

	// check handler-level attrs first
	for _, a := range h.attrs {
		switch a.Key {
		case "emoji":
			if s, ok := a.Value.Any().(string); ok {
				emoji = s
			}
		case "log_type":
			if s, ok := a.Value.Any().(string); ok {
				logType = s
			}
		case "no_newline":
			if b, ok := a.Value.Any().(bool); ok {
				noNewline = b
			}
		}
	}

	// check record attrs
	r.Attrs(func(a slog.Attr) bool {
		switch a.Key {
		case "emoji":
			if s, ok := a.Value.Any().(string); ok && emoji == "" {
				emoji = s
			}
		case "log_type":
			if s, ok := a.Value.Any().(string); ok && logType == "" {
				logType = s
			}
		case "no_newline":
			if b, ok := a.Value.Any().(bool); ok {
				noNewline = b
			}
		}
		return true
	})

	// if emoji not set, pick based on level
	if emoji == "" {
		if e, ok := h.LevelEmojis[r.Level]; ok {
			emoji = e
		} else {
			emoji = ""
		}
	}

	// choose color based on level and log_type
	var colorCode string
	switch r.Level {
	case slog.LevelInfo:
		colorCode = colorMagenta
		if logType == "input" {
			colorCode = colorReset
		}
		if logType == "success" {
			colorCode = colorGreen
		}
	case slog.LevelWarn:
		colorCode = colorYellow
	case slog.LevelError:
		colorCode = colorRed
	case slog.LevelDebug:
		colorCode = colorCyan
	default:
		colorCode = colorReset
	}

	// timestamp
	ts := time.Now()
	if !r.Time.IsZero() {
		ts = r.Time
	}
	timestamp := colorCode + ts.Format(h.TimeFormat) + colorReset

	// build message
	var b strings.Builder
	b.WriteString("[")
	b.WriteString(timestamp)
	b.WriteString("] ")
	if emoji != "" {
		b.WriteString(emoji)
	}
	b.WriteString(r.Message)

	// newline unless suppressed
	if !noNewline {
		b.WriteString("\n")
	}

	_, err := io.WriteString(h.Out, b.String())
	return err
}
