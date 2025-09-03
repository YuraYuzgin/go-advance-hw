package order

import (
	"cart-api/configs"
	"cart-api/pkg/middleware"
	"cart-api/pkg/req"
	"cart-api/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

type OrderHandler struct {
	OrderRepository *OrderRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
	}

	router.Handle("POST /order", middleware.IsAuth(handler.Create(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuth(handler.GetById(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuth(handler.GetAllByUserId(), deps.Config))
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[OrderRequest](&w, r)
		if err != nil {
			return
		}
		if len(body.Products) == 0 {
			http.Error(w, "Invalid products", http.StatusBadRequest)
			return
		}
		phone := r.Context().Value(middleware.ContextPhoneKey).(string)
		user, err := handler.OrderRepository.GetUserIdByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		productsInOrder, err := handler.OrderRepository.GetProductsById(body.Products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		createdOrder, err := handler.OrderRepository.Create(&Order{
			Model:    gorm.Model{},
			UserId:   user.ID,
			Products: *productsInOrder,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdOrder, 201)
	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		orderId, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid order ID format", http.StatusBadRequest)
			return
		}
		orderIdUint := uint(orderId)

		phone := r.Context().Value(middleware.ContextPhoneKey).(string)
		user, err := handler.OrderRepository.GetUserIdByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order, err := handler.OrderRepository.GetById(orderIdUint)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if user.ID != order.UserId {
			http.Error(w, "Invalid order ID", http.StatusBadRequest)
			return
		}
		res.Json(w, order, 200)
	}
}

func (handler *OrderHandler) GetAllByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone := r.Context().Value(middleware.ContextPhoneKey).(string)
		user, err := handler.OrderRepository.GetUserIdByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		orders, err := handler.OrderRepository.GetOrdersByUserId(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if len(*orders) == 0 {
			res.Json(w, nil, 200)
			return
		}
		res.Json(w, orders, 200)
	}
}
