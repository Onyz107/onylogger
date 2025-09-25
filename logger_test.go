package onylogger_test

import (
	"testing"

	"github.com/Onyz107/onylogger"
)

func TestOnylogger(t *testing.T) {
	log := onylogger.New()
	log.Debug("This is a debug message.")
	log.Info("This is an information message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")
	log.Success("This is a success message.")
	log.Fatal("This is a fatal message.")
}
