package repository

import (
	"go_test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoriRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoriRepositoryMock) FindById(id string) *entity.Category {
	argument := repository.Mock.Called(id)
	if argument.Get(0) == nil {
		return nil
	} else {
		result := argument.Get(0).(entity.Category)
		return &result
	}
}
