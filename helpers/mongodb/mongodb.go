package mongodb

import (
	"context"
	"log"

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
	Session    *mongo.Session
}

//NewMongoService returns an instance of MongoService
func NewMongoService(uri, name, collection string) *MongoService {
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(conext.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return &MongoService{
		DBName:     name,
		Collection: collection,
		Session:    client,
	}
}

// func (s *MongoService) InitiateRaid() error {

// }
