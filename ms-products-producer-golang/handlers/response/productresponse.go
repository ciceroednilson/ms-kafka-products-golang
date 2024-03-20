package response

import (
	"go/core/domain/productdomain"
)

// Product response
// @Description ProductResponse
// @Description Product fields
type ProductResponse struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price,omitempty"`
	Total       int16  `json:"total,omitempty"`
	Created     string `json:"created,omitempty"`
}

func NewListProductResponse(products []*productdomain.Product) []ProductResponse {
	total := len(products)
	var response []ProductResponse
	for i := 0; i < total; i++ {
		response = append(response, NewProductResponse(products[i]))
	}
	return response
}

func NewProductResponse(product *productdomain.Product) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Description: product.Description,
		Price:       product.Price,
		Total:       product.Total,
		Created:     product.Created,
	}
}
