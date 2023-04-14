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
	logger.SetLevel(log.WarnLevel)
	return logger
}
