package main

import (
	"github.com/lukasjarosch/belt/logging"
)

func main() {
	debug := true
	pretty := false
	log := logging.NewLogger(debug, pretty)

	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")

	// init logger via LOG_DEBUG and LOG_PRETTY
	logEnv := logging.NewFromEnvironment()
	logEnv.Debug("debug")
	logEnv.Info("info")
	logEnv.Warn("warning")
	logEnv.Error("error")
	logEnv.Fatal("fatal")
}
