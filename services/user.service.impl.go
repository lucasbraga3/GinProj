package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject.com/simple-api/models"
)

type UserServiceImpl struct {
	usercolleccion *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercolleccion.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUserbyName(name *string) (*models.User, error) {
	var user *models.User;
	query := bson.D{bson.E{Key:"name",Value: name}} //db.collections.find({name: "example"})
	err := u.usercolleccion.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(name *string, user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
