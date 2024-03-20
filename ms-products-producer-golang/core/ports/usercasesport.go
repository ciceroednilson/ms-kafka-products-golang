package ports

import "go/core/domain/productdomain"

type ProductUsercasePort interface {
	Save(product productdomain.Product) error
}
