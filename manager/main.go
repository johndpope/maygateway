package main

import (
	"github.com/gin-contrib/cors"
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
	router.Use(cors.Default())

	v1 := router.Group("v1")
	{
		apisGroup := v1.Group("apis")
		{
			apiController := new(controllers.ApiController)
			apisGroup.POST("/", apiController.Save)
			apisGroup.GET("/", apiController.GetAll)
			apisGroup.GET("/:id", apiController.GetOne)
			apisGroup.PUT("/:id", apiController.Update)
			apisGroup.DELETE("/:id", apiController.Delete)
		}
	}

	router.Run(":8080")

}
