package main

import (
	controller "github.com/kaburov38/GoShort/controller"
	database "github.com/kaburov38/GoShort/db"
	repository "github.com/kaburov38/GoShort/repository"
	service "github.com/kaburov38/GoShort/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)

	router := gin.Default()

	router.POST("/mapping", ctrl.Insert)
	router.GET("/read/:source", ctrl.Find)
	router.PUT("/mapping", ctrl.Update)
	router.DELETE("/mapping", ctrl.Delete)

	router.Run(":8080")
}
