package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject.com/simple-api/models"
)

type ProductServiceImpl struct {
	productcollection *mongo.Collection
	ctx               context.Context
}

func NewProductService(productcollection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{
		productcollection: productcollection,
		ctx:               ctx,
	}
}

func (p *ProductServiceImpl) CreateProduct(product *models.Product) error {
	_, err := p.productcollection.InsertOne(p.ctx, product)
	return err
}

func (p *ProductServiceImpl) GetProductbyid(id *int64) (*models.Product, error) {
	var product *models.Product
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := p.productcollection.FindOne(p.ctx, query).Decode(&product)
	return product, err
}

func (p *ProductServiceImpl) GetProductsbyName(name *string) ([]*models.Product, error) {
	var products []*models.Product
	query := bson.D{bson.E{Key: "name", Value: name}}
	cursor, err := p.productcollection.Find(p.ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductServiceImpl) GetProductsbyCategory(category *string) ([]*models.Product, error) {
	var products []*models.Product
	query := bson.D{bson.E{Key: "category", Value: category}}
	cursor, err := p.productcollection.Find(p.ctx, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductServiceImpl) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	cursor, err := p.productcollection.Find(p.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductServiceImpl) UpdateProduct(product *models.Product) error {
	filter := bson.D{bson.E{Key: "id", Value: product.Id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: product.Name}, bson.E{Key: "category", Value: product.Category},
		bson.E{Key: "description", Value: product.Description}, bson.E{Key: "price", Value: product.Price},
		bson.E{Key: "quantity", Value: product.Quantity}}}}
	result, _ := p.productcollection.UpdateOne(p.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("not found any doc to update")
	}
	return nil
}

func (p *ProductServiceImpl) DeleteProduct(productid *int64) error {
	filter := bson.D{bson.E{Key: "id", Value: productid}}
	result, _ := p.productcollection.DeleteOne(p.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("not found any doc to delete")
	}
	return nil
}
