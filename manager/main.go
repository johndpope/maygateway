package main

import (
	"github.com/gin-gonic/gin"
	"micro-gateway/manager/controllers"
	"micro-gateway/manager/db"
	"micro-gateway/manager/models"
)

func main() {
	db.Init()
	db.GetDB().AutoMigrate(&models.Api{})

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		apisGroup := v1.Group("apis")
		{
			apiController := new(controllers.ApiController)
			apisGroup.POST("/", apiController.Save)
		}
	}

	router.Run(":8080")

}
