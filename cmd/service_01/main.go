package main

import (
	"fmt"
	"os"

	"github.com/Jungle20m/golang-base/utils"
	"gorm.io/gorm"
)

// khởi tạo app context cần sử dụng cho app
type AppContext struct {
	mysql *gorm.DB
}

func main() {
	fmt.Println("main starting...")

	mysql, err := utils.GetMysqlConnection("anhnv:anhnv!@#456@tcp(1.53.252.177:3306)/healthnet?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		os.Exit(1)
	}

	appCtx := AppContext{
		mysql: mysql,
	}

	fmt.Printf("%+v\n", appCtx)
}
