package configs

import (
	"notee/models/users"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root@tcp(127.0.0.1:3306)/notee?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connection Database")
	}
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(&users.User{})
}
