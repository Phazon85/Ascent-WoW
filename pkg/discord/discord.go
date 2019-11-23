package discord

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

//Discord ...
type Discord struct {
	Token   string
	Keyword string
	DKP     DKP
}

//DKP ...
type DKP interface {
	InitRaidGroup(mc *discordgo.MessageCreate) error
}

//New ...
func New(dkp DKP, logger *zap.Logger, token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Debug("Failed to generate new Discord session", zap.String("Error", err.Error()))
	}

	//API Endpoints
	dg.AddHandler(MessageCreate)

	dg.Open()
	if err != nil {
		logger.Debug("Failed to start Discord event listeners", zap.String("Discord Reponse", err.Error()))
	}

	return dg
}
