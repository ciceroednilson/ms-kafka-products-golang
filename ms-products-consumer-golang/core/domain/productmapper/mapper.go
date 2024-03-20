package productmapper

import (
	"go/core/domain/productdomain"
	"go/infrastructure/entities"
	"time"
)

func ToEntity(p productdomain.Product) entities.Product {
	return entities.Product{
		Id:          p.Id,
		Description: p.Description,
		Price:       p.Price,
		Total:       p.Total,
		Created:     time.Now(),
	}
}
