package use_cases

import (
	"log"

	"github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"
)

// algo q seja privado ao pacote colocar minusculo
type createCategoryUseCase struct{
	// db instance
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

	return nil
}