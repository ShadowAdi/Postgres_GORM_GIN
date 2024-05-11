package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id    uint   `json:"ID" gorm:"primary_key"`
	Name  string `json:"username"`
	Email string `json:"email"`
	Books []Book `json:"books" gorm:"foreignkey:UserID"`
}

type Book struct {
	gorm.Model
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Author    string  `json:"author"`
	Price     float32 `json:"price"`
	UserID    uint    `json:"user_id"`
}
