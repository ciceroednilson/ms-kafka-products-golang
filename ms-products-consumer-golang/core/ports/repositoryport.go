package ports

import "module.mod/infrastructure/entities"

type ProductRepositoryPort interface {
	Save(product entities.Product) error
}
