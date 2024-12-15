package models

type User struct {
	Name    string `json:"name" bson:"name"`
	Phone   string `json:"phone" bson:"phone"`
	Email   string `json:"email" bson:"email"`
	Country string `json:"country" bson:"country"`
}
