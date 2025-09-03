package order

import (
	"cart-api/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId   uint
	Products []product.Product `gorm:"many2many:order_products;"`
}
