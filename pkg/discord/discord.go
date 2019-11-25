package discord

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

//Discord ...
type Discord struct {
	Logger *zap.Logger
	Prefix string
	
}

//DKP ...
type DKP interface {
	InitRaidGroup(mc *discordgo.MessageCreate) error
	StartRaid(mc *discordgo.MessageCreate) error
	StopRaid(mc *discordgo.MessageCreate) error
}

//New ...
func New(dkp DKP, logger *zap.Logger, token, prefix string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		logger.Debug("Failed to generate new Discord session", zap.String("Error", err.Error()))
	}

	d := &Discord{
		Logger: logger,
		Prefix: prefix,
	}

	//API Endpoints
	dg.AddHandler(messageCreate(dkp, d))
	dg.AddHandler(stateReady(prefix, logger))

	err = dg.Open()
	if err != nil {
		logger.Debug("Failed to start Discord event listeners", zap.String("Discord Reponse", err.Error()))
	}

	return dg
}
