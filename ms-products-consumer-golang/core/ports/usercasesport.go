package ports

import "module.mod/core/domain/productdomain"

type ProductUsecasePort interface {
	Save(product productdomain.Product) error
}
