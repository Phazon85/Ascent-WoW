package mongodb

import (
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
	Session    *mongo.MongoService
}

//NewMongoService returns an instance of MongoService
func NewMongoService(uri, name, collection string) *MongoService {
	session, err := mongo.NewClient(options.Client().ApplyURI(uri))

	return &MongoService{
		DBName:     name,
		Collection: collection,
		Session:    session,
	}
}

// func (s *MongoService) InitiateRaid() error {

// }
