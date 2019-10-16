package main

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/phazon85/Ascent-WoW/handlers"
	"github.com/phazon85/Ascent-WoW/helpers/logging"
)

//BotConfig holds system environment variables
type BotConfig struct {
	Token  string
	Config *handlers.Config
	Logger *logging.ZapLogger
}

// CheckBotConfig loads BOT_TOKEN and BOT_KEYWORD secrets
func CheckBotConfig() (*BotConfig, error) {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err := errors.New("BOT_TOKEN not defined in environment variables")
		return nil, err
	}

	log := logging.NewLogger()

	config := handlers.NewHandlers(os.Getenv("BOT_KEYWORD"), log)

	return &BotConfig{
		Token:  token,
		Config: config,
		Logger: log,
	}, nil
}

func main() {

	//load environment variables
	config, err := CheckBotConfig()
	if err != nil {
		config.Logger.Log.Debug("Failed to get required Bot Config")
	}
	defer config.Logger.Log.Sync()

	//Create new discordgo session
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		config.Logger.Log.Debug("Failed to generate new Discord session")
	}

	dg.AddHandler(config.Config.StateReady)
	dg.AddHandler(config.Config.MessageCreate)
	dg.AddHandler(config.Config.MessageReactionAdd)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		config.Logger.Log.Debug("Failed to start Discord event listeners")
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
