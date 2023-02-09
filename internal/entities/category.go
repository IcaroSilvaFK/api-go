package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func NewCategory(name string)( *Category, error){


	category := &Category{
		Name: name,
		CreateAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := category.IsValid()


	if(err != nil){
		return nil, err
	}
	// business rules

	return category,nil

}


func (category *Category) IsValid() error{

	if(len(category.Name) < 5){
		return fmt.Errorf("name must be grater than 5. got %d", len(category.Name))
	}
	return nil
}