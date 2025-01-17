package services

import "goproject.com/simple-api/models"

type ProductService interface {
	CreateProduct(*models.Product) error
	GetProductsbyName(*string) ([]*models.Product, error)
	GetProductbyid(*int64) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	GetProductsbyCategory(*string) ([]*models.Product, error)
	UpdateProduct(*models.Product) error
	DeleteProduct(*int64) error
}
