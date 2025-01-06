package controllers

import (
	"github.com/gin-gonic/gin"
	"goproject.com/simple-api/models"
	"goproject.com/simple-api/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productservice services.ProductService) ProductController {
	return ProductController{ProductService: productservice}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context)
func (pc *ProductController) GetProductbyid(*int64)
func (pc *ProductController) GetProductsbyName(*string)
func (pc *ProductController) GetProductsbyCategory(*string)
func (pc *ProductController) GetAll()
func (pc *ProductController) UpdateProduct(*models.Product)
func (pc *ProductController) DeleteProduct(*models.Product)
