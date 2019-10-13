package main

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/phazon85/Ascent-WoW/handlers"
)

//BotConfig holds system environment variables
type BotConfig struct {
	Token  string
	Config *handlers.Config
}

// CheckBotConfig loads BOT_TOKEN and BOT_KEYWORD secrets
func CheckBotConfig() (*BotConfig, error) {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err := errors.New("BOT_TOKEN not defined in environment variables")
		return nil, err
	}

	config := handlers.NewHandlers(os.Getenv("BOT_KEYWORD"))

	return &BotConfig{
		Token:  token,
		Config: config,
	}, nil
}

func main() {
	log.Println("Starting bot")

	//load environment variables
	config, err := CheckBotConfig()
	if err != nil {
		log.Printf("Error loading environment variables: %s", err.Error())
	}

	//Create new discordgo session
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Printf("Error creating new discord session: %s", err.Error())
	}

	dg.AddHandler(config.Config.StateReady)
	dg.AddHandler(config.Config.MessageCreate)
	dg.AddHandler(config.Config.MessageReactionAdd)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening discordgo session: %s", err.Error())
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
