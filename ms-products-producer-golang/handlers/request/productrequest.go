package request

import "go/core/domain/productdomain"

// Product request
// @Description ProductRequest
// @Description Product fields
type ProductRequest struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Total       int16  `json:"total"`
	Created     string `json:"created"`
}

func (p *ProductRequest) ToDomain() productdomain.Product {
	return productdomain.Product{
		Id:          p.Id,
		Description: p.Description,
		Price:       p.Price,
		Total:       p.Total,
		Created:     p.Created,
	}
}
