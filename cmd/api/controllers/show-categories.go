package controllers

import (
	"net/http"

	"github.com/IcaroSilvaFK/go-categories-msvc/internal/repositories"
	use_cases "github.com/IcaroSilvaFK/go-categories-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ShowCategories(ctx *gin.Context, repository repositories.ICategoryRepository)	{

	useCase := use_cases.NewListCategoriesUseCase(repository)

	categories,err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
		gin.H{
			"status":"Fail",
			"err":"Internal server error",
		})

		return
	}	


		ctx.JSON(http.StatusOK, gin.H{
			"categories":categories,
		})

}