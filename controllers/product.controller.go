package controllers

import (
	"net/http"
	"strconv"

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

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	reqerr := ctx.ShouldBindJSON(&product)
	if reqerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": reqerr.Error()})
		return
	}
	gaterr := pc.ProductService.CreateProduct(&product)
	if gaterr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": gaterr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "SUCCESS")
}

func (pc *ProductController) GetProductbyid(ctx *gin.Context) {
	productid, converr := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if converr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": converr.Error()})
		return
	}
	product, err := pc.ProductService.GetProductbyid(&productid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetProductsbyName(ctx *gin.Context) {
	name := ctx.Param("name")
	products, err := pc.ProductService.GetProductsbyName(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductsbyCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	products, err := pc.ProductService.GetProductsbyCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetAll(ctx *gin.Context) {
	products, err := pc.ProductService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	reqerr := ctx.ShouldBindJSON(&product)
	if reqerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": reqerr.Error()})
		return
	}
	gaterr := pc.ProductService.UpdateProduct(&product)
	if gaterr != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": gaterr.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "SUCCESS")
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	productid, converr := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if converr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": converr.Error()})
		return
	}
	err := pc.ProductService.DeleteProduct(&productid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "SUCCESS")
}

func (pc *ProductController) RegisterProductRoutes(router *gin.RouterGroup) {
	productroutes := router.Group("/products")
	productroutes.POST("/create", pc.CreateProduct)
	productroutes.GET("/all", pc.GetAll)
	productroutes.GET("/get/:id", pc.GetProductbyid)
	productroutes.GET("/getname/:name", pc.GetProductsbyName)
	productroutes.GET("/getcat/:category", pc.GetProductsbyCategory)
	productroutes.PATCH("/update", pc.UpdateProduct)
	productroutes.DELETE("/delete", pc.DeleteProduct)
}
