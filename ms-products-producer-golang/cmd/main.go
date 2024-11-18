package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"module.mod/core/usecases/productusecase"
	"module.mod/handlers/producthandler"
	"module.mod/infrastructure/repositories/productrepository"
)

// @title           Swagger Product API
// @version         1.0
// @description     This is a api to manager the products .
// @termsOfService  http://swagger.io/terms/

// @contact.name   Cícero Ednilson
// @contact.url    https://www.linkedin.com/in/c%C3%ADcero-ednilson-de-barros-machado-35153888/
// @contact.email  ciceroednilson@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000

// @externalDocs.description  Create Swagger docs with swaggo
// @externalDocs.url          https://github.com/swaggo/http-swagger
// @BasePath /
func main() {
	productrepository := productrepository.NewRepository()
	productusecase := productusecase.NewProductUseCase(productrepository)
	producthandler := producthandler.NewProductHandler(productusecase)
	initialServer(producthandler)
}

func initialServer(producthandler producthandler.ProductHandler) {
	router := mux.NewRouter()
	sub := router.PathPrefix("/products").Subrouter()
	sub.HandleFunc("", producthandler.Create).Methods(http.MethodPost)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         "localhost:7000",
	}

	log.Fatal(server.ListenAndServe())
}
