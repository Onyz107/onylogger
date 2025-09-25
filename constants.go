package onylogger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	colorReset   = "\033[0m"
	colorMagenta = "\033[35m"
	colorYellow  = "\033[33m"
	colorRed     = "\033[31m"
	colorCyan    = "\033[36m"
	colorGreen   = "\033[32m"
)

type OnyLogger struct {
	*logrus.Logger
}

type emojiFormatter struct {
	levelEmojis map[logrus.Level]string
}

func (f *emojiFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	emoji, ok := entry.Data["emoji"].(string)
	if !ok {
		emoji = f.levelEmojis[entry.Level]
	}

	var colorCode string
	switch entry.Level {
	case logrus.InfoLevel:
		colorCode = colorMagenta // Magenta for Info

		if logType, exists := entry.Data["log_type"].(string); exists && logType == "input" {
			colorCode = colorReset // No Color for Input
		}

		if logType, exists := entry.Data["log_type"].(string); exists && logType == "success" {
			colorCode = colorGreen
		}
	case logrus.WarnLevel:
		colorCode = colorYellow // Yellow
	case logrus.ErrorLevel:
		colorCode = colorRed // Red
	case logrus.DebugLevel:
		colorCode = colorCyan // Cyan
	default:
		colorCode = colorReset // Default (no color)
	}

	// Apply color to the timestamp
	timestamp := colorCode + entry.Time.Format("2006-01-02 15:04:05") + "\033[0m"

	var logMsg strings.Builder
	logMsg.WriteString("[")
	logMsg.WriteString(timestamp)
	logMsg.WriteString("] ")
	logMsg.WriteString(emoji)
	logMsg.WriteString(entry.Message)

	// Only add a newline if "no_newline" is not set to true.
	if noNewline, ok := entry.Data["no_newline"].(bool); !ok || !noNewline {
		logMsg.WriteString("\n")
	}

	return []byte(logMsg.String()), nil
}
