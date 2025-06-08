package repository

import (
    "context"
    "time"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
	"go-api/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type BlockRepository struct {
	Collection *mongo.Collection
}

func NewBlockRepository(db *mongo.Database) BlockRepository {
	return BlockRepository{
		Collection: db.Collection("block"),
	}
}


func (r BlockRepository) GetLastBlock() (*model.Block, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    var block model.Block
    opts := options.FindOne().SetSort(bson.D{{"timestamp", -1}})
    err := r.Collection.FindOne(ctx, bson.M{}, opts).Decode(&block)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil 
        }
        return nil, err 
    }

    return &block, nil
}

func (r BlockRepository) PostBlock(block model.Block) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Collection.InsertOne(ctx, block)
	if err != nil {
		return err
	}
	return nil
}


func (r BlockRepository) GetAllBlocks() ([]model.Block, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var blocks []model.Block
    for cursor.Next(ctx) {
        var block model.Block
        if err := cursor.Decode(&block); err != nil {
            return nil, err
        }
        blocks = append(blocks, block)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return blocks, nil
}