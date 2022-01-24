package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongoDB(collection *mongo.Collection) *MongoDB {
	return &MongoDB{collection: collection}
}

func (db *MongoDB) Create(ctx context.Context, item interface{}) (primitive.ObjectID, error) {
	res, err := db.collection.InsertOne(ctx, item)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// Boolean response here represents whether we created a new document
func (db *MongoDB) Update(ctx context.Context, item interface{}, idStr string) (bool, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false, err
	}
	res, err := db.collection.UpdateByID(ctx, id, bson.M{"$set": item}, options.Update().SetUpsert(true))
	if err != nil {
		return false, err
	}
	if res.MatchedCount == 0 {
		return true, nil
	}
	return false, nil
}

func (db *MongoDB) Get(ctx context.Context, idStr string, model interface{}) (bool, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false, err
	}
	if err := db.collection.FindOne(ctx, bson.M{"_id": id}).Decode(model); errors.Is(mongo.ErrNoDocuments, err) {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

func (db *MongoDB) List(ctx context.Context, models interface{}) error {
	cur, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	if err := cur.All(ctx, models); err != nil {
		return err
	}
	return nil
}

func (db *MongoDB) Delete(ctx context.Context, idStr string) (bool, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return false, err
	}
	res, err := db.collection.DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount == 0 {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
