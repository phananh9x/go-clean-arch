package mgo

import (
	"context"
	"go-clean-arch/service/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomersRepository interface {
	// Create a new customer
	Create(ctx context.Context, customer *entities.Customers) error

	// Retrieve a customer by their customer ID
	GetByID(ctx context.Context, customerID string) (*entities.Customers, error)

	// Update an existing customer
	Update(ctx context.Context, customer *entities.Customers) error

	// Delete an existing customer
	Delete(ctx context.Context, customerID string) error
}

type customersRepository struct {
	collection *mongo.Collection
}

func NewCustomersRepository(database *mongo.Database) CustomersRepository {
	collection := database.Collection("customers")
	return &customersRepository{collection}
}

func (repo *customersRepository) Create(ctx context.Context, customer *entities.Customers) error {
	_, err := repo.collection.InsertOne(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customersRepository) GetByID(ctx context.Context, customerID string) (*entities.Customers, error) {
	filter := bson.M{"customer_id": customerID}
	var result entities.Customers
	err := repo.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *customersRepository) Update(ctx context.Context, customer *entities.Customers) error {
	filter := bson.M{"customer_id": customer.CustomerId}
	update := bson.M{"$set": customer}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (repo *customersRepository) Delete(ctx context.Context, customerID string) error {
	filter := bson.M{"customer_id": customerID}
	_, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
