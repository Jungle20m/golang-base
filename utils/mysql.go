package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetMysqlConnection(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
dbConn.DB()
dbConn.DB().Ping()
dbConn.DB().SetMaxIdleConns(10)
dbConn.DB().SetMaxOpenConns(100)
*/
