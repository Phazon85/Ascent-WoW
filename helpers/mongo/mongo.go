package mongo

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/mgo.v2"
)

type DKPService interface {
}

//MongoService implements DKPService
type MongoService struct {
	DBName     string
	Collection string
	Session    *mgo.Session
}

//NewMongoService returns an instance of MongoService
func NewMongoService(uri, name, collection string) *MongoService {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoService{
		DBName:     name,
		Collection: collection,
		Session:    session,
	}
}

//InitiateRaid starts recording of members in raid for DKP
func (m *MongoService) InitiateRaid(msg *discordgo.MessageCreate) {

}
