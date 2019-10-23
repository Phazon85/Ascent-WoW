package main

import (
	"errors"
	"os"

	"go.uber.org/zap"

	"github.com/bwmarrin/discordgo"
	"github.com/phazon85/Ascent-WoW/handlers"
	"github.com/phazon85/Ascent-WoW/helpers/logging"
)

const (
	postgresFile = "dev.yaml"
	driverName   = "postgres"
)

//BotConfig holds system environment variables
type BotConfig struct {
	Token  string
	Config *handlers.Config
	Logger *logging.ZapLogger
}

// CheckBotConfig loads BOT_TOKEN and BOT_KEYWORD secrets
func CheckBotConfig() (*BotConfig, error) {

	log := logging.NewLogger()

	token := os.Getenv("BOT_TOKEN")
	keyword := os.Getenv("BOT_KEYWORD")
	if token == "" || keyword == "" {
		err := errors.New("Failed to load environment variable")
		return nil, err
	}

	sql := postgres.NewDBObject(postgresFile, driverName)

	config := handlers.NewHandlers(keyword, log)

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
		cfg.Logger.Log.Debug("Failed to get required Bot Config", zap.String("Error", err.Error()))
	}
	defer cfg.Logger.Log.Sync()

	//Create new discordgo session
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		cfg.Logger.Log.Debug("Failed to generate new Discord session", zap.String("Error", err.Error()))
	}

	dg.AddHandler(cfg.Config.StateReady)
	dg.AddHandler(cfg.Config.MessageCreate)
	dg.AddHandler(cfg.Config.MessageReactionAdd)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		cfg.Logger.Log.Debug("Failed to start Discord event listeners", zap.String("Discord Reponse", err.Error()))
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
