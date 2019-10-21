package main

import (
	"errors"
	"os"

	"github.com/phazon85/Ascent-WoW/helpers/mongo"

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

	uri := os.Getenv("MONGO_URI")
	collection := os.Getenv("MONGO_COLLECTION")
	name := os.Getenv("MONGO_NAME")

	token := os.Getenv("BOT_TOKEN")
	keyword := os.Getenv("BOT_KEYWORD")
	if token == "" || keyword == "" {
		err := errors.New("Variables not defined in environment")
		return nil, err
	}

	log := logging.NewLogger()

	mongo := mongo.NewMongoService(uri, name, collection)

	config := handlers.NewHandlers(keyword, log, mongo)

	return &BotConfig{
		Token:  token,
		Config: config,
		Logger: log,
	}, nil
}

func main() {

	//load environment variables
	cfg, err := CheckBotConfig()
	if err != nil {
		cfg.Logger.Log.Debug("Failed to get required Bot Config")
	}
	defer cfg.Logger.Log.Sync()

	//Create new discordgo session
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		cfg.Logger.Log.Debug("Failed to generate new Discord session")
	}

	dg.AddHandler(cfg.Config.StateReady)
	dg.AddHandler(cfg.Config.MessageCreate)
	dg.AddHandler(cfg.Config.MessageReactionAdd)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		cfg.Logger.Log.Debug("Failed to start Discord event listeners")
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
