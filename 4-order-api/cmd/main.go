package main

import (
	"cart-api/configs"
	"cart-api/internal/auth"
	"cart-api/internal/product"
	"cart-api/pkg/db"
	"cart-api/pkg/middleware"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Repositories
	productRepository := product.NewProductRepository(db)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	// Middlewares
	stack := middleware.Chain(middleware.Logging)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}
