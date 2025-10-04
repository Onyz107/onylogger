package onylogger

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func (l *OnyLogger) Input(message string, userInput *string) {
	l.logger.Info(message,
		slog.String("log_type", "input"),
		slog.String("emoji", "[📝] "),
		slog.Bool("no_newline", true),
	)

	fmt.Print(" ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*userInput = scanner.Text()
}

func (l *OnyLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, args...))
}

func (l *OnyLogger) Debug(message string) {
	l.logger.Debug(message)
}

func (l *OnyLogger) Infof(format string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, args...))
}

func (l *OnyLogger) Info(message string) {
	l.logger.Info(message)
}

func (l *OnyLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *OnyLogger) Warn(message string) {
	l.logger.Warn(message)
}

func (l *OnyLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, args...))
}

func (l *OnyLogger) Error(message string) {
	l.logger.Error(message)
}

func (l *OnyLogger) Successf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.logger.Info(msg,
		slog.String("log_type", "success"),
		slog.String("emoji", "[✅] "),
	)
}

func (l *OnyLogger) Success(message string) {
	l.Successf("%s", message)
}
