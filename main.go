package main

import (
	"gorm-test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "test:test@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Human{})
	db.AutoMigrate(&models.Account{})
}
