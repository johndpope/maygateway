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

func (api ApiController) GetAll(ctx *gin.Context) {
	apis, err := api.model.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &apis)
}

func (api ApiController) GetOne(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	apiSearch, err := api.model.GetOne(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &apiSearch)
}

func (api ApiController) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var apiModel models.Api
	if err := ctx.BindJSON(&apiModel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	udpatedApi, err := api.model.Update(id, apiModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, udpatedApi)
}

func (api ApiController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	err := api.model.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)

}