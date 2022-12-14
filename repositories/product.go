package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(Id int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Preload("Category").Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(Id int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").Preload("Category").First(&product, Id).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}
