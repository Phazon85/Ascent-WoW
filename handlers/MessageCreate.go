package handlers

import (
	"log"
	"strings"

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
