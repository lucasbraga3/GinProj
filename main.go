package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"goproject.com/simple-api/controllers"
	"goproject.com/simple-api/services"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercolleccion *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongo")
}

func main() {

}
