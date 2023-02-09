package use_cases

import (
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/repositories"
)

type listCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}


func NewListCategoriesUseCase(repository repositories.ICategoryRepository) * listCategoriesUseCase {
	return &listCategoriesUseCase{repository}
}

func (useCase *listCategoriesUseCase) Execute() ([]*entities.Category,error) {

	categories, err := useCase.repository.Show()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
