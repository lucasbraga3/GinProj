package models

type Product struct {
	Id          int64  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Price       int    `json:"price" bson:"price"`
	Quantity    int    `json:"quantity" bson:"quantity"`
	Category    string `json:"category" bson:"category"`
}
