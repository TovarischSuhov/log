package main

import "github.com/TovarischSuhov/log"

func main() {
	log.Debug("Message")
	log.Info("Message")
	log.Warn("Message")
	log.Error("Message")
	log.UseColors = true
	log.Debug("Colored Message")
	log.Info("Colored Message")
	log.Warn("Colored Message")
	log.Error("Colored Message")
}
