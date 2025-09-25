package onylogger

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func New() *OnyLogger {
	log := logrus.New()
	log.SetFormatter(&emojiFormatter{
		levelEmojis: map[logrus.Level]string{
			logrus.InfoLevel:  "[📜 ] ",
			logrus.WarnLevel:  "[⚠️ ] ",
			logrus.ErrorLevel: "[ ❌ ] ",
			logrus.DebugLevel: "[ 🐛 ] ",
		},
	})
	return &OnyLogger{Logger: log}
}

func (l *OnyLogger) Input(message string, userInput *string) {
	l.WithField("log_type", "input").
		WithField("emoji", "[📝] ").
		WithField("no_newline", true).
		Info(message)
	fmt.Print(" ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*userInput = scanner.Text()
}

func (l *OnyLogger) Success(message string) {
	l.WithField("log_type", "success").
		WithField("emoji", "[ ✅ ] ").
		Info(message)
}

func (l *OnyLogger) Successf(format string, args ...interface{}) {
	l.WithField("log_type", "success").
		WithField("emoji", "[ ✅ ]").
		Infof(format, args...)
}
