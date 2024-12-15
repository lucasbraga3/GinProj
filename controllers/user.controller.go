package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goproject.com/simple-api/models"
	"goproject.com/simple-api/services"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{UserService: userservice}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User;
	reqerr := ctx.ShouldBindBodyWithJSON(&user); 
	if reqerr!= nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message":reqerr.Error()});
		return;
	}
	gateerr := uc.UserService.CreateUser(&user);
	if gateerr != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message":gateerr.Error()});
		return;
	}

	ctx.JSON(http.StatusOK, "SUCCESS")
}

func (uc *UserController) GetUserbyName(ctx *gin.Context) {
	username := ctx.Param("name");
	user, err := uc.UserService.GetUserbyName(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message":err.Error()});
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(router *gin.RouterGroup) {
	userroutes := router.Group("/user")
	userroutes.POST("/create", uc.CreateUser)
	userroutes.GET("/get/:name", uc.GetUserbyName)
	userroutes.GET("/getall", uc.GetAll)
	userroutes.PATCH("/update", uc.UpdateUser)
	userroutes.DELETE("/delete", uc.DeleteUser)
}
