package handlers

import (
	"github.com/bwmarrin/discordgo"
)

//Actions implements the config struct
type Actions interface {
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
	MessageReactionAdd(s *discordgo.Session, mra *discordgo.MessageReactionAdd)
	StateReady(s *discordgo.Session, r *discordgo.Ready)
}

//Config implements the Config interface
type Config struct {
	BotKeyword string
}

//NewHandlers takes in helper and info
func NewHandlers(
	botkey string,
) *Config {
	return &Config{
		BotKeyword: botkey,
	}
}
