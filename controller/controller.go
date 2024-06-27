package controller

import (
	modelApi "github.com/kaburov38/GoShort/model/api"
	"github.com/kaburov38/GoShort/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Insert(ctx *gin.Context) {
	var request modelApi.InsertRequest
	ctx.BindJSON(&request)

	err := c.service.Insert(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (c *Controller) Find(ctx *gin.Context) {
	var request modelApi.FindRequest
	request.Source = ctx.Param("source")

	response, err := c.service.Find(request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}

func (c *Controller) Update(ctx *gin.Context) {
	var request modelApi.UpdateRequest
	ctx.BindJSON(&request)

	err := c.service.Update(request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

func (c *Controller) Delete(ctx *gin.Context) {
	var request modelApi.DeleteRequest
	ctx.BindJSON(&request)

	err := c.service.Delete(request)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
