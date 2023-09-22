package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:mysql@tcp(127.0.0.1:3306)/go-credential?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	DB = database
}
