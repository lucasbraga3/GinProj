package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject.com/simple-api/models"
)

type UserServiceImpl struct {
	usercolleccion *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercolleccion *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercolleccion: usercolleccion,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercolleccion.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUserbyName(name *string) (*models.User, error) {
	var user *models.User;
	query := bson.D{bson.E{Key: "name", Value: name}}; //db.collections.find({name: "example"})
	err := u.usercolleccion.FindOne(u.ctx, query).Decode(&user);
	return user, err;
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User;
	cursor, err := u.usercolleccion.Find(u.ctx, bson.D{{}});
	if err != nil {
		return nil, err;
	}
	for cursor.Next(u.ctx) {
		var user models.User;
		err := cursor.Decode(&user);
		if err != nil {
			return nil, err;
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "email", Value: user.Email}};
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "phone", Value: user.Phone}, bson.E{Key: "country", Value: user.Country}}}};
	result, _ := u.usercolleccion.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("not found any doc to update")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(email *string) error {
	filter := bson.D{bson.E{Key: "email", Value: email}}
	result, _ := u.usercolleccion.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("not found any doc to update")
	}
	return nil
}
