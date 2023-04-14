package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper  *viper.Viper
	Db     *gorm.DB
	Logger *logrus.Logger
)
