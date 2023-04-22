package repository

import "go_test/entity"

type CategoriRepository interface {
	FindById(id string) *entity.Category
}
