package initialize

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() *log.Logger {
	//set json for type of log
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	file, _ := os.Create("runtime/info.log")
	logger.SetOutput(file)
	logger.WithError(errors.New("123")).Info("do")
	logger.SetLevel(log.WarnLevel)
	return logger
}
