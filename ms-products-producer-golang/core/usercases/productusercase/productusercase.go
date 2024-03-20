package productusercase

import (
	"go/core/domain/productdomain"
	"go/core/domain/productmapper"
	"go/core/ports"
	"log"
)

type ProductUserCase struct {
	ProductRepositoryPort ports.ProductRepositoryPort
}

func NewProductUserCase(productRepositoryPort ports.ProductRepositoryPort) ports.ProductUsercasePort {
	return &ProductUserCase{
		ProductRepositoryPort: productRepositoryPort,
	}
}

func (p *ProductUserCase) Save(product productdomain.Product) error {
	log.Println("Starting the save a product.")
	entity := productmapper.ToEntity(product)
	return p.ProductRepositoryPort.Save(entity)
}
