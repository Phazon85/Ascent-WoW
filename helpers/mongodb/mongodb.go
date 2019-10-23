package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//DKPService contains actions for interacting with the MongoDB client
type DKPService interface {
	InitiateRaid() error
	LoadRaids()
}

//RaidGroup holds guild, past raid information, and current DKP standings
type RaidGroup struct {
	Guild string
	Raids []raids
	DKP   map[string]int
}

type raids struct {
	Name  string
	Group map[string]bool
	Date  string
}

// MongoService implements DKPService
type MongoService struct {
	DBName     string
	Collection string
	Session    *mongo.Client
}

//NewMongoService returns an instance of MongoService
func NewMongoService(uri, name, col string) (*MongoService, error) {
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

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoService{
		DBName:     name,
		Collection: col,
		Session:    client,
	}, nil
}

//LoadRaids loads all known raids into a...
func (s *MongoService) LoadRaids() []RaidGroup {
	groups := []RaidGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := s.Session.Database(s.DBName).Collection(s.Collection).Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem RaidGroup
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		groups = append(groups, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return groups
}

// //InitiateRaid will place raid in ready state for members to join
// func (s *MongoService) InitiateRaid(channelid) error {
// 	_, err := s.Session.Database(s.DBName).Collection(s.Collection).InsertOne(context.Background(), bson.M{
// 		"Name":    "Ascent",
// 		"Members": "Test",
// 		"Date":    time.Now(),
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
