package onylogger

import (
	"OnyLogger/internal/handlers"
	"io"
	"log/slog"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type OnyLogger struct {
	logger   *slog.Logger
	levelVar *slog.LevelVar
}

func New(out io.Writer) *OnyLogger {
	levelVar := new(slog.LevelVar)
	levelVar.Set(LevelInfo)

	levelEmojis := map[slog.Level]string{
		LevelDebug: "[🐛] ",
		LevelInfo:  "[📜] ",
		LevelWarn:  "[⚠️] ",
		LevelError: "[❌] ",
	}

	handler := &handlers.EmojiHandler{
		Out:         out,
		LevelEmojis: levelEmojis,
		TimeFormat:  "2006-01-02 15:04:05",
		LevelVar:    levelVar,
	}

	logger := slog.New(handler)
	return &OnyLogger{logger: logger, levelVar: levelVar}
}

func (l *OnyLogger) SetLevel(level slog.Level) {
	l.levelVar.Set(level)
}
