package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func colourPrint() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {

	logfile, _ := os.OpenFile("Randoms/example.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	multiWriter := io.MultiWriter(os.Stdout, logfile) // To write to stdout and logfile at the same time

	loggerInstance := log.New()
	loggerInstance.SetOutput(multiWriter)
	loggerInstance.SetFormatter(&log.JSONFormatter{})

	loggerInstance.Trace("Something very low level.")
	loggerInstance.Debug("Useful debugging information.")
	loggerInstance.Info("Something noteworthy happened!")
	loggerInstance.Warn("You should probably take a look at this.")
	loggerInstance.Error("Something failed but I'm not quitting.")
	loggerInstance.WithFields(log.Fields{"animal": "walrus", "size": 10}).Info("A group of walrus emerges from the ocean")
	loggerInstance.WithFields(log.Fields{"omg": true, "number": 122}).Warn("The group's number increased tremendously!")

}
