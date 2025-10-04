package onylogger_test

import (
	"OnyLogger"
	"bytes"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := onylogger.New(&buf)

	if logger == nil {
		t.Fatal("Expected logger instance, got nil")
	}

	// Test logging output
	logger.Info("test message")
	out := buf.String()
	if !strings.Contains(out, "test message") || !strings.Contains(out, "[📜]") {
		t.Fatalf("Expected log output to contain message and emoji, got: %s", out)
	}
}

func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := onylogger.New(&buf)

	logger.SetLevel(slog.LevelDebug)

	// Test that debug message now logs
	logger.Debug("debug message")
	out := buf.String()
	if !strings.Contains(out, "debug message") || !strings.Contains(out, "[🐛]") {
		t.Fatalf("Expected debug log output with emoji, got: %s", out)
	}
}

func TestOutput(t *testing.T) {
	logger := onylogger.New(os.Stderr)
	logger.SetLevel(slog.LevelDebug)
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.Success("success message")
}
