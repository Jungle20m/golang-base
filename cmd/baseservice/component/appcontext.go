package component

import (
	"gorm.io/gorm"
)

type appCtx struct {
	database *Database
}

func NewAppContext(db *Database) *appCtx {
	return &appCtx{database: db}
}

func (ctx *appCtx) GetMysqlConnection() *gorm.DB {
	return ctx.database.mysql
}
