package repositories

import "github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func  NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category,0),
	}
}	

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {

	r.db = append(r.db,category)

	return nil
}	

func (r *inMemoryCategoryRepository) Show() ([]*entities.Category,error){

	return  r.db, nil
}