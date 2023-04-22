package service

import (
	"go_test/entity"
	"go_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoriRepositoryMock = &repository.CategoriRepositoryMock{Mock: mock.Mock{}}
var CategoryService = CategoriService{Repository: CategoriRepositoryMock}

func TestCategoryService_Get(t *testing.T) {
	CategoriRepositoryMock.Mock.On("FindById", "2").Return(nil)
	category, err := CategoryService.Get("2")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetCategory(t *testing.T) {
	result := entity.Category{
		Id:   "2",
		Name: "Kresna",
	}
	CategoriRepositoryMock.Mock.On("FindById", "2").Return(result)
	category, err := CategoryService.Get("2")
	assert.NotNil(t, category)
	assert.Nil(t, err)
}
