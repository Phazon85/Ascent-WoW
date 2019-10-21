package handlers

import (
	"log"
	"strings"

	"go.uber.org/zap/zapcore"

	"github.com/bwmarrin/discordgo"
)

//MessageCreate is the callback for the MessageCrate event from Discord
func (c *Config) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	//Checks if bot keyword is present in message
	if !strings.HasPrefix(m.Content, c.BotKeyword) {
		return
	}

	// [BOT_KEYWORD] [command] [options] :: "!asc" "log_level" "debug"
	contentTokens := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(contentTokens) < 2 {
		return
	}

	switch strings.ToLower(contentTokens[1]) {
	case "info":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, infoMessage())
		if err != nil {
			log.Printf("Error sending info message: %s", err.Error())
			return
		}

	case "help":
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, helpMessage())
		if err != nil {
			log.Printf("Error sending help message: %s", err.Error())
			return
		}
	case "log_level":
		if len(contentTokens) >= 3 {
			levelOpt := strings.ToLower(contentTokens[2])

			switch levelOpt {
			case "debug":
				c.Logger.Level.SetLevel(zapcore.DebugLevel)
				c.Logger.Log.Debug("Log level changed to Debug")
			case "info":
				c.Logger.Level.SetLevel(zapcore.InfoLevel)
				c.Logger.Log.Info("Log level changed to Info")
			case "warn":
				c.Logger.Level.SetLevel(zapcore.WarnLevel)
				c.Logger.Log.Warn("Log level changed to Warn")
			case "error":
				c.Logger.Level.SetLevel(zapcore.ErrorLevel)
				c.Logger.Log.Error("Log level changed to Error")
			case "fatal":
				c.Logger.Level.SetLevel(zapcore.FatalLevel)
			case "panic":
				c.Logger.Level.SetLevel(zapcore.PanicLevel)
			}
		}

	case "raid":
		if len(contentTokens) >= 3 {
			raidOpt := strings.ToLower(contentTokens[2])

			switch raidOpt {
			case "start":
				c.Logger.Log.Debug("Starting Raid")
				// c.Mongo.InitiateRaid()
				s.ChannelMessageSend(m.ChannelID, "Raid has been started")
			}

		}
	default:
		//Silently fail for any unregistered commands
	}
}

func infoMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
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
}

func helpMessage() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
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
}
