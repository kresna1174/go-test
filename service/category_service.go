package service

import (
	"errors"
	"go_test/entity"
	"go_test/repository"
)

type CategoriService struct {
	Repository repository.CategoriRepository
}

func (service CategoriService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Data Not Found")
	} else {
		return category, nil
	}
}
