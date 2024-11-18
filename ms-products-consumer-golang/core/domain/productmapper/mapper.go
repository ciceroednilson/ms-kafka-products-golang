package productmapper

import (
	"time"

	"module.mod/core/domain/productdomain"
	"module.mod/infrastructure/entities"
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
