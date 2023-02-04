package repositories

import "github.com/IcaroSilvaFK/go-categories-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
}