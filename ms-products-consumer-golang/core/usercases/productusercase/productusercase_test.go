package productusercase_test

import (
	"go/core/domain/productdomain"
	"go/core/ports"
	"go/core/usercases/productusercase"
	"go/infrastructure/entities"
	"testing"

	"github.com/stretchr/testify/assert"
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
	productusercase := productusercase.NewProductUserCase(productrepositoryMock)
	product := productdomain.Product{
		Description: "Bola",
		Price:       1.00,
		Total:       10,
		Created:     "2024-01-05 02:58:06",
	}
	err := productusercase.Save(product)
	assert.Nil(t, err)
}
