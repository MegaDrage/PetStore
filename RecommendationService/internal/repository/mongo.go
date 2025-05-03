package repository

import (
	"context"

	"MegaDrage/PetStore/RecommendationService/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository(uri string) (*MongoRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoRepository{client: client}, nil
}

func (r *MongoRepository) SaveRecommendation(ctx context.Context, rec models.Recommendation) error {
	collection := r.client.Database("recommendations").Collection("recs")
	_, err := collection.InsertOne(ctx, rec)
	return err
}