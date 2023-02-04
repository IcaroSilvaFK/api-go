package controllers

import (
	"net/http"

	use_cases "github.com/IcaroSilvaFK/go-categories-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput  struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context){

	var body createCategoryInput;

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
		gin.H{
			"success":false,
			"err": err.Error(),
		})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase()

	err := useCase.Execute((body.Name))

	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
				"success": false,
				"error": err.Error(),
			})

			return 
	}

	ctx.JSON(http.StatusCreated, gin.H{"success":true})

}