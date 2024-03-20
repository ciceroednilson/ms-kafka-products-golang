package productmapper

import (
	"fmt"
	"go/core/domain/productdomain"
	"go/infrastructure/entities"
	"strconv"
	"time"
)

func EntityToDomain(product *entities.Product) *productdomain.Product {
	return &productdomain.Product{
		Id:          product.Id,
		Description: product.Description,
	}
}

func EntityToDomainFull(product *entities.Product) *productdomain.Product {
	price := fmt.Sprintf("%.2f", product.Price)
	return &productdomain.Product{
		Id:          product.Id,
		Description: product.Description,
		Price:       price,
		Total:       product.Total,
		Created:     product.Created.Format("2006-01-02 15:04:05"),
	}
}

func ToItemsDomain(products []entities.Product) []*productdomain.Product {
	list := []*productdomain.Product{}
	for _, product := range products {
		domain := EntityToDomainFull(&product)
		list = append(list, domain)
	}
	return list
}

func ToEntity(p productdomain.Product) entities.Product {
	price, _ := strconv.ParseFloat(p.Price, 2)
	return entities.Product{
		Id:          p.Id,
		Description: p.Description,
		Price:       price,
		Total:       p.Total,
		Created:     time.Now(),
	}
}
