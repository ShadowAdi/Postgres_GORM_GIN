package main

import (
	"github.com/ShadowAdi/gin_gorm_postgres/config"
	"github.com/ShadowAdi/gin_gorm_postgres/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
