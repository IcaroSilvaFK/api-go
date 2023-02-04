package use_cases

import (
	"log"

	"github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"
	"github.com/IcaroSilvaFK/go-categories-msvc/internal/repositories"
)

// algo q seja privado ao pacote colocar minusculo
type createCategoryUseCase struct{
	repository repositories.ICategoryRepository
}


func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (useCase *createCategoryUseCase) Execute(name string) error{

	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: persist entity to db
	log.Fatalln(category)
	err = useCase.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}