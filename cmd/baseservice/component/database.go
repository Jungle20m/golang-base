package component

import (
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	mysql *gorm.DB
}

func (db *Database) Mysql() *gorm.DB {
	return db.mysql
}

func NewDatabase(config *Config) (*Database, error) {
	log.Info("Connecting to mysql server")
	db, err := gorm.Open(mysql.Open(config.Mysql.DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Info("Connected to mysql server successfully")
	return &Database{
		mysql: db,
	}, err
}
