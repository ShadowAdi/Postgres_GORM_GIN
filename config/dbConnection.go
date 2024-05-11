package config

import (
	"github.com/ShadowAdi/gin_gorm_postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	urlConnect := "host=localhost user=postgres password=Aditya00 dbname=NewData port=5432"
	db, err := gorm.Open(postgres.Open(urlConnect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Book{})
	DB = db
}
