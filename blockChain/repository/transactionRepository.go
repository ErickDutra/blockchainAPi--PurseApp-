package repository

import (
	"context"
	"go-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type TransactionRepository struct {
	Collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) TransactionRepository {
	return TransactionRepository{
		Collection: db.Collection("transaction"),
	}
}

func (repo TransactionRepository) GetAllTransactions() ([]model.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []model.Transaction
	for cursor.Next(ctx) {
		var transaction model.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo TransactionRepository) PostTransaction(transaction model.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.Collection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}


func (repo TransactionRepository) GetTransactionByID(id string) (*model.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var transaction model.Transaction
	err := repo.Collection.FindOne(ctx, bson.M{"id": id}).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil 
		}
		return nil, err 
	}

	return &transaction, nil
}

func (r TransactionRepository) UpdateStatusAndSignature(id string, signature string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    filter := bson.M{"id": id}
    update := bson.M{"$set": bson.M{"status": model.StatusConfirmed, "signature": signature}}
    _, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}
