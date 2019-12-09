package main

import (
	"log"
	"os"

	"github.com/phazon85/ascent-wow/pkg/discord"
	"go.uber.org/zap"
)

const (
	//TOKEN ...
	TOKEN = "ASCENT_WOW_TOKEN"
	//KEYWORD ...
	KEYWORD = "ASCENT_WOW_KEYWORD"
	//ENV ...
	ENV = "ASCENT_WOW_ENV"
	//DEFAULTENV ...
	DEFAULTENV = "dev"
)

func main() {

	//load environment
	token := os.Getenv(TOKEN)
	keyword := os.Getenv(KEYWORD)
	// dbconn := os.Getenv(DBCONN)
	if token == "" || keyword == "" {
		log.Fatal("Connection and token strings required")
	}

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = DEFAULTPORT
	// }

	env := os.Getenv(ENV)
	if env == "" {
		env = DEFAULTENV
	}

	var logger *zap.Logger
	var err error
	switch env {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatal(err)
	}

	// db := multisql.NewSQL(dbconn)
	// pgsql := pgsql.New(db)
	// dkp := dkp.New(pgsql)

	dg := discord.New(logger, token, keyword)
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
