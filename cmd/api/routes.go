package main

import (
	"github.com/IcaroSilvaFK/go-categories-msvc/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(ginInstance *gin.Engine) {

	router := ginInstance.Group("/_v1") 

	router.POST("/category",controllers.CreateCategory)
}