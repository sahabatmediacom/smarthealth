package config

import (
	"fmt"
	"pamer-api/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta")

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Hospital{})
	db.AutoMigrate(&entity.Paramedic{})
	db.AutoMigrate(&entity.User{})

	DB = db
}
