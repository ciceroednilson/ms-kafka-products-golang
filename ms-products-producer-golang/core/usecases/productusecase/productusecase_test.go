package productusecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"module.mod/core/domain/productdomain"
	"module.mod/core/ports"
	"module.mod/core/usecases/productusecase"
	"module.mod/infrastructure/entities"
)

type ProductRepositoryMock struct{}

func NewProductRepositoryMock() ports.ProductRepositoryPort {
	return &ProductRepositoryMock{}
}

func (p *ProductRepositoryMock) Save(product entities.Product) error {
	return nil
}

func Test_SaveProduct(t *testing.T) {
	productrepositoryMock := NewProductRepositoryMock()
	productusecase := productusecase.NewProductUseCase(productrepositoryMock)
	product := productdomain.Product{
		Description: "Bola",
		Price:       "1.00",
		Total:       10,
		Created:     "2024-01-05 02:58:06",
	}
	err := productusecase.Save(product)
	assert.Nil(t, err)
}
