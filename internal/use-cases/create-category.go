package use_cases

import (
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/repositories"
)

// algo q seja privado ao pacote colocar minusculo
type createCategoryUseCase struct{
	repository repositories.ICategoryRepository
}


func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (useCase *createCategoryUseCase) Execute(name string) error{

	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: verify if category name already exists
	err = useCase.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}