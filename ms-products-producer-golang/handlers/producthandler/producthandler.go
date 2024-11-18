package producthandler

import (
	"encoding/json"

	"net/http"

	"module.mod/core/ports"
	"module.mod/handlers/request"
)

type ProductHandler struct {
	ProductUsecasePort ports.ProductUsecasePort
}

func NewProductHandler(productUsecasePort ports.ProductUsecasePort) ProductHandler {
	return ProductHandler{
		ProductUsecasePort: productUsecasePort,
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
	err = p.ProductUsecasePort.Save(domain)
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
