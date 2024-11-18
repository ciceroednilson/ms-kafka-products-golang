package productusecase

import (
	"log"

	"module.mod/core/domain/productdomain"
	"module.mod/core/domain/productmapper"
	"module.mod/core/ports"
)

type ProductUseCase struct {
	ProductRepositoryPort ports.ProductRepositoryPort
}

func NewProductUseCase(productRepositoryPort ports.ProductRepositoryPort) ports.ProductUsecasePort {
	return &ProductUseCase{
		ProductRepositoryPort: productRepositoryPort,
	}
}

func (p *ProductUseCase) Save(product productdomain.Product) error {
	log.Println("Starting the save a product.")
	entity := productmapper.ToEntity(product)
	return p.ProductRepositoryPort.Save(entity)
}
