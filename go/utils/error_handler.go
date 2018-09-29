package utils

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	log.SetFormatter(&formatter)

	logFile, _ := os.OpenFile("log.json", os.O_WRONLY|os.O_CREATE, 0755)
	log.SetOutput(logFile)
	log.SetLevel(log.InfoLevel)
}

func HandleError(err error, msg string, exit bool) {
	if err != nil {
		logMsg := log.WithFields(log.Fields{
			"error": err,
		})

		if exit {
			logMsg.Fatal(msg)
		}

		logMsg.Error(msg)
	}
}

func HandleCommandError(err error, out, msg string, exit bool) {
	if err != nil {
		logMsg := log.WithFields(log.Fields{
			"error": err,
			"out":   out,
		})

		if exit {
			logMsg.Fatal(msg)
		}

		logMsg.Error(msg)
	}
}

func Log(obj interface{}, msg string) {
	log.WithFields(log.Fields{
		"Object": obj,
	}).Info(msg)
}
