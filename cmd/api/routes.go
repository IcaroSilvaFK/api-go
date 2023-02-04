package main

import (
	"github.com/IcaroSilvaFK/go-categories-msvc/cmd/api/controllers"
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(ginInstance *gin.Engine) {

	router := ginInstance.Group("/_v1") 

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	router.POST("/category",func (ctx *gin.Context) {
		controllers.CreateCategory(ctx,	inMemoryCategoryRepository)
	})

	router.GET("/categories",func (ctx *gin.Context) {
		controllers.ShowCategories(ctx,	inMemoryCategoryRepository)
	})
}