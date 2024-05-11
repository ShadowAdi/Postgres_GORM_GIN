package routes

import (
	"github.com/ShadowAdi/gin_gorm_postgres/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
	router.POST("/", controller.PostController)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
	router.GET("/:id", controller.GetOne)

}
