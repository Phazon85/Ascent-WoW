package main

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/phazon85/Ascent-WoW/services"
)

//BotConfig holds system environment variables
type BotConfig struct {
	Token  string
	Config *services.Config
}

// CheckBotConfig loads BOT_TOKEN and BOT_KEYWORD secrets
func CheckBotConfig() (*BotConfig, error) {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err := errors.New("BOT_TOKEN not defined in environment variables")
		return nil, err
	}

	helpMessage := &discordgo.MessageEmbed{
		URL:   "https://github.com/Phazon85/Ascent-WoW",
		Title: "Help Section",
		Color: 0x031cfc,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Help Section",
				Value: `
			help - Displays this! :D
			info - Displays information about the bot
			`,
				Inline: false,
			},
		},
	}

	infoMessage := &discordgo.MessageEmbed{
		URL:    "https://github.com/Phazon85/Ascent-WoW",
		Title:  "Ascent-WoW Discord Bot",
		Color:  0x031cfc,
		Footer: &discordgo.MessageEmbedFooter{Text: "Made using the discordgo library"},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "About",
				Value:  "Ascent-WoW is a bot... To be added",
				Inline: false,
			},
			{
				Name:   "Author",
				Value:  "Ascent-WoW bot was created by Phazon85/Tourian",
				Inline: false,
			},
			{
				Name:   "Library",
				Value:  "Ascent-WoW uses the discordgo library by bwmarrin",
				Inline: false,
			},
		},
	}

	config := services.NewHandlers(helpMessage, infoMessage, os.Getenv("BOT_KEYWORD"))

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

	dg.AddHandler(callbackConfig.StateReady)
	dg.AddHandler(callbackConfig.MessageCreate)

	//Starts discord event listener
	err = dg.Open()
	if err != nil {
		log.Printf("Error opening discordgo session: %s", err.Error())
	}
	defer dg.Close()

	//Keeps listeners alive
	<-make(chan struct{})

}
