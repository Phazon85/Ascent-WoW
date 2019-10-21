package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type DKPService interface {
// 	InitiateRaid() error
// }

// MongoService implements ContactService
type MongoService struct {
	DBName     string
	Collection string
	Session    *mongo.Client
}

//NewMongoService returns an instance of MongoService
func NewMongoService(uri, name, collection string) (*MongoService, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &MongoService{
		DBName:     name,
		Collection: collection,
		Session:    client,
	}, nil
}

// func (s *MongoService) InitiateRaid() error {

// }
