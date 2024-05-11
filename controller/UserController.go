package controller

import (
	"net/http"

	"github.com/ShadowAdi/gin_gorm_postgres/config"
	"github.com/ShadowAdi/gin_gorm_postgres/model"
	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context) {
	Users := []model.User{}
	config.DB.Preload("Books").Find(&Users)

	c.JSON(200, &Users)
}

func PostController(c *gin.Context) {
	var userInput struct {
		model.User
		Books []model.Book `json:"books"`
	}
	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&userInput.User)
	for _, book := range userInput.Books {
		book.UserID = userInput.User.Id
		config.DB.Create(&book)
	}
	c.JSON(http.StatusCreated, userInput)
}

func Delete(c *gin.Context) {
	var user model.User
	config.DB.Where("id=?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func Update(c *gin.Context) {
	var user model.User
	userID := c.Param("id")

	result := config.DB.Preload("Books").First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Association("books").Clear()

	config.DB.Save(&user)

	if len(user.Books) > 0 {
		for _, book := range user.Books {
			book.UserID = user.ID // Set the user ID for each book
			config.DB.Save(&book)
		}
	}

	c.JSON(http.StatusOK, &user)

}

func GetOne(c *gin.Context) {
	var user model.User
	config.DB.Preload("Books").Where("id=?", c.Param("id")).First(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User Get", "data": user})
}
