package initialize

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() *log.Logger {
	//set json for type of log
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	_, err := os.Create("runtime/info.log")
	logger.WithError(err).Info("do")
	logger.SetLevel(log.WarnLevel)
	return logger
}
