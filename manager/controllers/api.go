package controllers

import (
	"github.com/gin-gonic/gin"
	"micro-gateway/manager/models"
	"net/http"
)

type ApiController struct {
	model models.ApiModel
}

func (api ApiController) Save(context *gin.Context) {
	var apiModel models.Api

	if err := context.BindJSON(&apiModel); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	_, err := api.model.Save(apiModel)
	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.Status(http.StatusCreated)

	return
}