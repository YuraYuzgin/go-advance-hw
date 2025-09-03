package main

import (
	"cart-api/configs"
	"cart-api/internal/auth"
	"cart-api/internal/order"
	"cart-api/internal/product"
	"cart-api/internal/user"
	"cart-api/pkg/db"
	"cart-api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)
	orderRepository := order.NewOrderRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            conf,
	})
	order.NewOrderHandler(router, order.OrderHandlerDeps{
		OrderRepository: orderRepository,
		Config:          conf,
	})

	// Middlewares
	stack := middleware.Chain(middleware.Logging)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Serve is listening on port 8081")
	server.ListenAndServe()
}
