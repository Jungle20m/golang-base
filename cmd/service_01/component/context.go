package component

import (
	"github.com/Jungle20m/golang-base/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMysqlConnection() *gorm.DB
	GetKafkaInstance() *utils.KafkaInstance
	GetConfig() *Config
	GetLoggerInstance() *logrus.Logger
}

type appCtx struct {
	db     *gorm.DB
	kaf    *utils.KafkaInstance
	conf   *Config
	logger *logrus.Logger
}

func NewAppContext(db *gorm.DB, kafka *utils.KafkaInstance, config *Config, logger *logrus.Logger) *appCtx {
	return &appCtx{
		db:     db,
		kaf:    kafka,
		conf:   config,
		logger: logger,
	}
}

func (ctx *appCtx) GetMysqlConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetKafkaInstance() *utils.KafkaInstance {
	return ctx.kaf
}

func (ctx *appCtx) GetConfig() *Config {
	return ctx.conf
}

func (ctx *appCtx) GetLoggerInstance() *logrus.Logger {
	return ctx.logger
}
