package productconsumer

import (
	"go/consumer/request"
	"go/core/ports"
)

type ProductConsumer struct {
	ProductUsercasePort ports.ProductUsercasePort
}

func NewProductConsumer(productUsercasePort ports.ProductUsercasePort) ProductConsumer {
	return ProductConsumer{
		ProductUsercasePort: productUsercasePort,
	}
}

func (p *ProductConsumer) Create(request request.ProductRequest) {
	domain := request.ToDomain()
	err := p.ProductUsercasePort.Save(domain)
	if err != nil {
		panic(err)
	}
}
