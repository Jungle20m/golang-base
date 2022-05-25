package component

import (
	"gorm.io/gorm"
)

type AppCtx struct {
	database *Database
}

func NewAppContext(db *Database) *AppCtx {
	return &AppCtx{database: db}
}

func (ctx *AppCtx) GetMysqlConnection() *gorm.DB {
	return ctx.database.mysql
}
