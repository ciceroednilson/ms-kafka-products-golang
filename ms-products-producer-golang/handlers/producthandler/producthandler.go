package producthandler

import (
	"encoding/json"
	"go/core/ports"
	"go/handlers/request"
	"net/http"
)

type ProductHandler struct {
	ProductUsercasePort ports.ProductUsercasePort
}

func NewProductHandler(productUsercasePort ports.ProductUsercasePort) ProductHandler {
	return ProductHandler{
		ProductUsercasePort: productUsercasePort,
	}
}

// Create a product
// @Summary      Product
// @Description  create register of product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param request body request.ProductRequest true "query params"
// @Success      200
// @Router       /products [post]
func (p *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request request.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	domain := request.ToDomain()
	err = p.ProductUsercasePort.Save(domain)
	if err != nil {
		panic(err)
	}
	responseSuccessfully(w, nil)
}

func responseSuccessfully(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	if body != nil {
		w.Write(body)
	}
	w.WriteHeader(http.StatusOK)
}
