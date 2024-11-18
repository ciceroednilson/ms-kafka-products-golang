package productconsumer

import (
	"module.mod/consumer/request"
	"module.mod/core/ports"
)

type ProductConsumer struct {
	ProductUsecasePort ports.ProductUsecasePort
}

func NewProductConsumer(productUsecasePort ports.ProductUsecasePort) ProductConsumer {
	return ProductConsumer{
		ProductUsecasePort: productUsecasePort,
	}
}

func (p *ProductConsumer) Create(request request.ProductRequest) {
	domain := request.ToDomain()
	err := p.ProductUsecasePort.Save(domain)
	if err != nil {
		panic(err)
	}
}
