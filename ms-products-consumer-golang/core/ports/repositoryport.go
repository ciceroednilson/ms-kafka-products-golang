package ports

import (
	"go/infrastructure/entities"
)

type ProductRepositoryPort interface {
	Save(product entities.Product) error
}
