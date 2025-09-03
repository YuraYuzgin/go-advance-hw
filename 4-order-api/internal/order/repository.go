package order

import (
	"cart-api/internal/product"
	"cart-api/internal/user"
	"cart-api/pkg/db"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	result := repo.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repo *OrderRepository) GetUserIdByPhone(phone string) (*user.User, error) {
	var user user.User
	result := repo.Database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *OrderRepository) GetProductsById(ids []uint) (*[]product.Product, error) {
	var products []product.Product
	result := repo.Database.DB.
		Table("products").
		Where("id IN ?", ids).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (repo *OrderRepository) GetById(id uint) (*Order, error) {
	var order Order
	result := repo.Database.DB.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepository) GetOrdersByUserId(id uint) (*[]Order, error) {
	var orders []Order
	result := repo.Database.DB.
		Table("orders").
		Where("user_id = ?", id).
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}
