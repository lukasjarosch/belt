package main

import "github.com/lukasjarosch/belt/logging"

func main() {
	debug := true
	pretty := false
	log := logging.NewLogger(debug, pretty)

	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	log.Fatal("fatal")
}
