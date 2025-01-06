package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"goproject.com/simple-api/controllers"
	"goproject.com/simple-api/services"
)

var (
	server            *gin.Engine
	userservice       services.UserService
	usercontroller    controllers.UserController
	productservice    services.ProductService
	productcontroller controllers.ProductController
	ctx               context.Context
	usercolleccion    *mongo.Collection
	productcollection *mongo.Collection
	mongoclient       *mongo.Client
	err               error
)

func init() {
	ctx = context.TODO()
	err := godotenv.Load("mongouser.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	//mongodb://user:password@localhost:0000
	credential := options.Credential{
		Username:   os.Getenv("USER"),
		Password:   os.Getenv("PWD"),
		AuthSource: os.Getenv("DB"),
	}
	dbname := os.Getenv("DB")
	uri := "mongodb://" + os.Getenv("HOST") + ":" + os.Getenv("PORT")
	mongoconn := options.Client().ApplyURI(uri).SetAuth(credential)
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongo")

	usercolleccion = mongoclient.Database(dbname).Collection("users")
	userservice = services.NewUserService(usercolleccion, ctx)
	usercontroller = controllers.New(userservice)
	productcollection = mongoclient.Database(dbname).Collection("products")
	productservice = services.NewProductService(productcollection, ctx)
	productcontroller = controllers.NewProductController(productservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/api")
	usercontroller.RegisterUserRoutes(basepath)
	productcontroller.RegisterProductRoutes(basepath)
	log.Fatal(server.Run(":8080"))
}
