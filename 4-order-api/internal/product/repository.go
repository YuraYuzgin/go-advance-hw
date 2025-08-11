package product

import (
	"cart-api/pkg/db"
	"errors"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) (int, error) {
	result := repo.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return 500, result.Error
	}
	if result.RowsAffected == 0 {
		return 404, errors.New("link not found")
	}
	return 200, nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() (*[]Product, error) {
	var products []Product
	result := repo.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}
