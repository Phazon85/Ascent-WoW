package main

import (
	"log"
	"os"

	"github.com/phazon85/Ascent-WoW/pkg/discord"
	"github.com/phazon85/Ascent-WoW/pkg/dkp"
	"github.com/phazon85/Ascent-WoW/pkg/repository/pgsql"
	"github.com/phazon85/multisql"
	"go.uber.org/zap"
)

const (
	//DBCONN ...
	DBCONN = "ASCENT_WOW_DBCONN"
	//TOKEN ...
	TOKEN = "ASCENT_WOW_TOKEN"
	//KEYWORD ...
	KEYWORD = "ASCENT_WOW_KEYWORD"
	//PORT ...
	PORT = "ASCENT_WOW_PORT"
	//DEFAULTPORT ...
	DEFAULTPORT = "9000"
	//ENV ...
	ENV = "ASCENT_WOW_ENV"
	//DEFAULTENV ...
	DEFAULTENV = "dev"
)

func main() {

	//load environment
	token := os.Getenv(TOKEN)
	keyword := os.Getenv(KEYWORD)
	dbconn := os.Getenv(DBCONN)
	if token == "" || keyword == "" || dbconn == "" {
		log.Fatal("Connection and token strings required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULTPORT
	}

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

	db := multisql.NewSQL(dbconn)
	pgsql := pgsql.New(db)
	dkp := dkp.New(pgsql)

	dg := discord.New(dkp, logger, token, keyword)
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
